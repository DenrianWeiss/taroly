package auth

import (
	"github.com/DenrianWeiss/taroly/service/db"
	"log"
	"os"
)

const AUTH_USER = "auth_user_"
const ROOT_USER = "root"

func init() {
	uid, present := os.LookupEnv("TAROLY_ROOT_USER")
	if !present {
		panic("TAROLY_ROOT_USER not set")
	}
	err := db.Set(db.GetDb(), []byte(AUTH_USER+uid), []byte(""))
	if err != nil {
		log.Panicln("init auth error: ", err.Error())
	}
	err = db.Set(db.GetDb(), []byte(ROOT_USER), []byte(uid))
	if err != nil {
		log.Panicln("init auth error: ", err.Error())
	}
}

func IsRoot(uid string) bool {
	root, err := db.Get(db.GetDb(), []byte(ROOT_USER))
	if err != nil {
		return false
	}
	return string(root) == uid
}

func IsAuth(uid string) bool {
	result, err := db.Exist(db.GetDb(), []byte(AUTH_USER+uid))
	if err != nil {
		return false
	}
	return result
}

func AddAuth(uid string) {
	db.Set(db.GetDb(), []byte(AUTH_USER+uid), []byte(""))
}
