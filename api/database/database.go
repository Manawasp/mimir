package database

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"

	"mimir/api/config"
	"mimir/api/utils"
)

// Session TODO
var Session *mgo.Session

func init() {
	var err error
	Session, err = New()
	if err != nil {
		log.Errorf("Unable to connect to the database: %v.", err)
		return
	}
}

// New TODO
func New() (*mgo.Session, error) {
	conf := config.App.DB

	dialInfo, err := mgo.ParseURL(conf.Host)
	if err != nil {
		log.Errorf("Unable to connect the datastore: %v", err)
		return nil, err
	}
	dialInfo.Database = conf.Database
	dialInfo.Timeout = time.Second * 10

	if len(conf.Username) > 0 {
		dialInfo.Username = conf.Username
		dialInfo.Password = conf.Password
	}

	if conf.SSL {
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{
				InsecureSkipVerify: true,
			})
		}
	}

	session, err := mgo.DialWithInfo(dialInfo)
	// session, err := mgo.Dial(conf.Host)
	if err != nil {
		log.Errorf("Unable to connect the datastore: %v", err)
		return nil, err
	}

	if err := session.Ping(); err != nil {
		return nil, err
	}

	return session, nil
}

// ToContext TODO
func ToContext(ctx context.Context, db *mgo.Database) context.Context {
	newctx := context.WithValue(ctx, utils.KeyDB, db)
	return newctx
}

// FromContext TODO
func FromContext(ctx context.Context) *mgo.Database {
	if db := ctx.Value(utils.KeyDB); db != nil {
		return db.(*mgo.Database)
	}
	return nil
}
