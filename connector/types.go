package connector

import (
	"database/sql"
)

type AuthOption struct {
	Url      string
	UserName string
	Password string
}

type ConnectHandle struct {
	DB      *sql.DB
	Session *sql.Conn
}

type LogAuth struct {
	Auth  AuthOption
	Conn  map[int]*ConnectHandle
	State string
}

type LogMinterConnector struct {
	Auth     [7]*LogAuth // 最大支持7个线程
	Parallel int8        // 线程并发数, 默认4个
	Threads  int8        // 当前建立的线程
}
