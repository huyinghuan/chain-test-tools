/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package birdsnest serialize
package birdsnest

import (
	"encoding/json"
	"reflect"
	"time"

	birdsnestpb "chainmaker.org/chainmaker/common/v2/birdsnest/pb"

	"github.com/gogo/protobuf/proto"
)

// TODO Split BirdsNestImpl and Serialize

// Start serialize monitor TODO Goroutinue should be turned off using context.Context here
func (b *BirdsNestImpl) Start() {
	go b.serializeMonitor()
	go b.serializeTimed()
}

// serializeMonitor
func (b *BirdsNestImpl) serializeMonitor() {
	for { // nolint
		select {
		// Only signals for the current filter "serialized type" are received
		case signal := <-b.serializeC:
			t, ok := SerializeIntervalType_name[signal.typ]
			if !ok {
				b.log.Errorf("serialize type %v not support", t)
			}
			switch signal.typ {
			case SerializeIntervalType_Height:
				// current height - pre height < height interval does not serialize; otherwise, it serialize
				// eg: 85 - 80 = 5 < 10
				// 	   5 < 10 true does not serialize
				if b.height-b.preHeight.Load() < b.config.Snapshot.BlockHeight.Interval {
					continue
				}
			case SerializeIntervalType_Timed, SerializeIntervalType_Exit:
				// common.SerializeIntervalType_Timed and common.SerializeIntervalType_Exit are handled directly
			default:
				continue
			}
			err := b.Serialize()
			if err != nil {
				b.log.Errorf("serialize error type: %v, error: %v", t, err)
			}
		}
	}
}

// Serialize all cuckoos in the current BirdsNest
func (b *BirdsNestImpl) Serialize() error {
	t := time.Now()
	// Logs are print after the method is complete
	defer func(log Logger) {
		elapsed := time.Since(t)
		log.Debugf("bird's nest serialize success elapsed: %v", elapsed)
	}(b.log)
	// convert []CuckooFilter to []*common.CuckooFilter
	var filters []*birdsnestpb.CuckooFilter
	filters, err := analysisCuckooFilters(b.filters)
	if err != nil {
		return err
	}
	marshal, err := json.Marshal(b.config)
	if err != nil {
		return err
	}
	birdsNest := &birdsnestpb.BirdsNest{
		Config:       marshal,
		Height:       b.preHeight.Load(),
		CurrentIndex: uint32(b.currentIndex),
		Filters:      filters,
	}
	data, err := proto.Marshal(birdsNest)
	if err != nil {
		return err
	}
	// Write to disk
	err = b.snapshot.Write(data)
	if err != nil {
		return err
	}
	// Increase the height
	b.preHeight.Store(b.height)
	return nil
}

// Deserialize deserialize Bird's nest
func (b *BirdsNestImpl) Deserialize() error {
	data, err := b.snapshot.Read()
	if err != nil {
		return err
	}
	if data == nil {
		return nil
	}
	var bn = new(birdsnestpb.BirdsNest)
	err = proto.Unmarshal(data, bn)
	if err != nil {
		return err
	}
	filters, err := newCuckooFiltersByDecode(bn.Filters)
	if err != nil {
		return err
	}
	var bnConfig BirdsNestConfig
	err = json.Unmarshal(bn.Config, &bnConfig)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(bnConfig, b.config) {
		err = ErrCannotModifyTheNestConfiguration
	}
	b.filters = filters
	b.config = &bnConfig
	b.height = bn.Height
	b.currentIndex = int(bn.CurrentIndex)
	return err
}

// serializeTimed send SerializeIntervalType_Timed sign
func (b *BirdsNestImpl) serializeTimed() {
	if b.config.Snapshot.Type != SerializeIntervalType_Timed {
		return
	}
	ticker := time.NewTicker(time.Second * time.Duration(b.config.Snapshot.Timed.Interval))
	// nolint
	for {
		select {
		case <-ticker.C:
			b.serializeC <- serializeSignal{typ: SerializeIntervalType_Timed}
		}
	}
}

// serializeExit send SerializeIntervalType_Exit sign
// nolint: unused
func (b *BirdsNestImpl) serializeExit() {
	b.serializeC <- serializeSignal{typ: SerializeIntervalType_Exit}
}

// serializeHeight send SerializeIntervalType_Height sign
func (b *BirdsNestImpl) serializeHeight(height uint64) {
	if b.config.Snapshot.Type != SerializeIntervalType_Height {
		return
	}
	b.serializeC <- serializeSignal{typ: SerializeIntervalType_Height, height: height}
}

// Serialize signal
type serializeSignal struct {
	typ    SerializeIntervalType
	height uint64
}
