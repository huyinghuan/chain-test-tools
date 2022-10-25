package services

import (
	"chain-api-imgo/utils"
	"sync"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

var aliveLock = sync.RWMutex{}
var alivePool = make(map[string]*sdk.ChainClient)

func GetKeepAliveChainClient(cacheClientToPool bool, userKeyBytes, userCertBytes, userSignKeyBytes, userSignCertBytes []byte, args ...string) (chainClient *sdk.ChainClient, fromCache bool, err error) {
	addr, _, _, err := utils.ParseAddrAndSkiFromCrtBytes(userCertBytes)
	if err == nil {
		aliveLock.RLock()
		if c, ok := alivePool[addr]; ok {
			fromCache = true
			aliveLock.RUnlock()
			return c, fromCache, nil
		}
		aliveLock.RUnlock()
	}
	client, err := GetUserChainClient(userKeyBytes, userCertBytes, userSignKeyBytes, userSignCertBytes, args...)
	if err != nil {
		return nil, fromCache, err
	}
	if cacheClientToPool {
		aliveLock.Lock()
		alivePool[addr] = client
		aliveLock.Unlock()
		fromCache = true
	}
	return client, fromCache, nil
}
