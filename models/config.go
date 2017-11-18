package models

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/itokun12/face-maker/orm"
	"github.com/itokun12/face-maker/utils"
	"github.com/pkg/errors"
)

var (
	masterDB *orm.DB
	logger   = logrus.New()
)

type ConnectionType int

const (
	_ ConnectionType = iota
	MySQL
)

func Initialize() {
	masterDB = connectToDB(MySQL, utils.Env("DATABASE_SCHEMA", "face_maker"))
}

func Finalize() {
	masterDB.Close()
}

func connectToDB(typ ConnectionType, schemaName string) *orm.DB {
	var db *orm.DB
	var err error
	if typ == MySQL {
		db, err = orm.NewDBFromEnvValue(schemaName)
	} else {
		e := fmt.Errorf("Unknown connection type %v", typ)
		logger.Errorln(e)
		panic(e)
	}

	if err != nil {
		e := errors.Wrapf(err, "Connect to database failed (schemaName:%v)", schemaName)
		logger.Errorln(e)
		panic(e)
	}
	return db
}
