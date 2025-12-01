package connection

import "sync"

type conn struct {
	creds map[string]string
}

func (c *conn) Sett() {
	//Setter Function
}

func (c *conn) Get() {
	//Getter Function
}

var (
	c    *conn
	once sync.Once
)

func NewConn() *conn {
	once.Do(func() {
		c = &conn{
			creds: make(map[string]string),
		}
	})

	return c
}
