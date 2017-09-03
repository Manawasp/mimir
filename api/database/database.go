package database

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"

	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"

	"feedback/api/config"
)

var Session *mgo.Session

// New TODO
func New() error {
	conf := config.App.DB

	if len(conf.Host) == 0 {
		return errors.New("Database is not configured")
	}

	// Generate connection string
	// "mongodb://<username>:<password>@<hostname>:<port>,<hostname>:<port>/<db-name>
	str := "mongodb://"
	if len(conf.Username) > 0 {
		str += conf.Username + ":" + conf.Password + "@"
	}
	str += conf.Host + "/" + conf.Database
	dialInfo, err := mgo.ParseURL(str)

	// Setup TLS connection if HTTPS set to true
	if conf.HTTPS {
		roots := x509.NewCertPool()
		roots.AppendCertsFromPEM([]byte(""))

		tlsConfig := &tls.Config{}
		tlsConfig.RootCAs = roots

		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	Session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Errorf("Unable to connect the datastore: %v", err)
		return err
	}

	if err := Session.Ping(); err != nil {
		return err
	}

	return nil
}
