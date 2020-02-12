package main

import (
	"github.com/go-pg/migrations"
)

var schema = `

CREATE TABLE conn_log (
	user_id bigint,
	ip_addr varchar(15),
	ts timestamp
);
`

func init() {
	migrations.Register(func(db migrations.DB) error {
		_, err := db.Exec(schema)
		if err != nil {
			return err
		}

		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			DROP TABLE if exists conn_log;
			DROP TABLE if exists user_id_to_ip`)
		if err != nil {
			return err
		}

		return err
	})
}
