package user

import (
	"fmt"
	"github.com/DenrianWeiss/taroly/model"
	"github.com/DenrianWeiss/taroly/service/cache/rpc"
	"sync"
)

var userStatusMap = sync.Map{}

func GetUserStatus(uid string) model.UserStatus {
	if status, ok := userStatusMap.Load(uid); ok {
		return status.(model.UserStatus)
	} else {
		userStatusMap.Store(uid, model.UserStatus{})
	}
	return model.UserStatus{}
}

func GetUserAccount(uid string) string {
	return GetUserStatus(uid).CurrentAccount
}

func GetUserOnlineMode(uid string) bool {
	return GetUserStatus(uid).OnlineMode
}

func SetUserAccount(uid string, account string) {
	status := GetUserStatus(uid)
	status.CurrentAccount = account
	userStatusMap.Store(uid, status)
}

func SetUserOnlineMode(uid string, onlineMode bool) {
	status := GetUserStatus(uid)
	status.OnlineMode = onlineMode
	userStatusMap.Store(uid, status)
}

func SetUserSimulation(uid string, pid int, port int) {
	status := GetUserStatus(uid)
	status.SimulationPid = pid
	status.SimulationPort = port
	userStatusMap.Store(uid, status)
}

func SetUserChain(uid string, chain string) {
	status := GetUserStatus(uid)
	status.Chain = chain
	userStatusMap.Store(uid, status)
}

func GetUserRpcUrl(uid string) string {
	s := GetUserStatus(uid)
	if s.OnlineMode {
		chain := s.Chain
		if chain == "" {
			return ""
		} else {
			return rpc.GetRpcUrl(chain)
		}
	} else {
		port := s.SimulationPort
		if port == 0 {
			return ""
		} else {
			return fmt.Sprintf("http://127.0.0.1:%d", port)
		}
	}
	return ""
}
