/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package birdsnest snapshot
package birdsnest

import (
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"chainmaker.org/chainmaker/common/v2/wal"
)

// WalSnapshot wal snapshot
type WalSnapshot struct {
	snapshotM sync.Mutex
	wal       *wal.Log
}

// NewWalSnapshot new wal snapshot eg: data/tx_filter/chainN/birdsnestN
func NewWalSnapshot(path, name string, number int) (*WalSnapshot, error) {
	opts := wal.DefaultOptions
	opts.NoSync = false
	if number > 0 {
		// eg: data/txfilter/chainN/birdnest1
		path = filepath.Join(path, name+strconv.Itoa(number))
	} else {
		// eg: data/txfilter/chainN/birdnest
		path = filepath.Join(path, name)
	}
	err := createDirIfNotExist(path)
	if err != nil {
		return nil, err
	}
	file, err := wal.Open(path, opts)
	if err != nil {
		return nil, err
	}
	ws := &WalSnapshot{
		wal:       file,
		snapshotM: sync.Mutex{},
	}
	return ws, nil
}

// Read safe read wal
func (s *WalSnapshot) Read() ([]byte, error) {
	s.snapshotM.Lock()
	defer s.snapshotM.Unlock()
	index, err := s.wal.LastIndex()
	if err != nil {
		return nil, err
	}
	if index == 0 {
		return nil, nil
	}
	read, err := s.wal.Read(index)
	if err != nil {
		return nil, err
	}
	return read, nil
}

// Write safe write wal
func (s *WalSnapshot) Write(data []byte) error {
	s.snapshotM.Lock()
	defer s.snapshotM.Unlock()
	index, err := s.wal.LastIndex()
	if err != nil {
		return err
	}
	index++
	err = s.wal.Write(index, data)
	if err != nil {
		return err
	}
	err = s.wal.TruncateFront(index)
	if err != nil {
		return err
	}
	return nil
}

// createDirIfNotExist create dir if not exist
func createDirIfNotExist(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		// create dir
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

//type LwsSnapshot struct {
//	snapshotM sync.Mutex
//	lws       *lws.Lws
//}
//
//func NewLwsSnapshot(path, name string, number int) (*LwsSnapshot, error) {
//	if number > 0 {
//		path = filepath.Join(path, name+strconv.Itoa(number))
//	} else {
//		path = filepath.Join(path, name)
//	}
//	l, err := lws.Open(path,
//		lws.WithWriteFlag(lws.WF_TIMEDFLUSH, 1000),
//		lws.WithEntryLimitForPurge(2),
//		//lws.WithWriteFlag()
//	)
//	if err != nil {
//		return nil, err
//	}
//	//l.
//	err = createDirIfNotExist(path)
//	if err != nil {
//		return nil, err
//	}
//	ws := &LwsSnapshot{
//		lws:       l,
//		snapshotM: sync.Mutex{},
//	}
//	return ws, nil
//}
//
//func (s *LwsSnapshot) Read() ([]byte, error) {
//
//	return nil, nil
//}
//
//func (s *LwsSnapshot) Write(data []byte) error {
//	return nil
//}
