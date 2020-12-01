package utils

import "strconv"

// GetDSN func
func GetDSN(instance string, user string, pass string, host string, port int, database string, ssl bool) (dsn string) {
	dsn = instance + "://"
	if user != "" && pass != "" {
		dsn = dsn + user + ":" + pass + "@"
	}
	sslMode := "disable"
	if ssl {
		sslMode = "enable"
	}
	dsn = dsn + host + ":" + strconv.Itoa(port) + "/" + database + "?sslmode=" + sslMode
	return
}
