package connector

import (
	"database/sql"
	"github.com/godror/godror"
	"time"
)

const (
	DefaultTimeZone = "UTC"
	PoolConnectMin  = 1
	PoolConnectMax  = 1
)

func (c *LogMinterConnector) SetDefault(a AuthOption, tz string) {

	var cs godror.ConnectionParams
	cs.ConnectString = a.Url
	cs.Username = a.UserName
	cs.Password = godror.NewPassword(a.Password)
	cs.SessionTimeout = 3
	local, _ := time.LoadLocation(DefaultTimeZone) // "Asia/Shanghai"
	cs.Timezone = local
	cs.SetSessionParamOnInit("NLS_NUMERIC_CHARACTERS", ",.")
	cs.SetSessionParamOnInit("NLS_LANGUAGE", "AMERICAN")

	conn := sql.OpenDB(godror.NewConnector(cs))
	conn.SetMaxIdleConns(PoolConnectMin)
	conn.SetMaxOpenConns(PoolConnectMax)
}

func NewConnect(...[]AuthOption) {

}
