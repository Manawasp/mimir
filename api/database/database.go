package database

import (
	"crypto/tls"
	"crypto/x509"
	"net"

	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"

	"feedback/api/config"
)

// New TODO
func New() (*mgo.Session, error) {
	conf := config.AppConfig.DB

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

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Errorf("Unable to connect the datastore: %v", err)
		return nil, err
	}

	if err := session.Ping(); err != nil {
		return nil, err
	}

	return session, nil
}
