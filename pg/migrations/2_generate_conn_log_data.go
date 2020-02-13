package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("Generating test data. Please wait...")

		_, err := db.Exec(`
			INSERT INTO conn_log (user_id, ip_addr, ts) VALUES
				(1, '10.10.0.1', generate_series('2019-01-01 00:00'::timestamp, '2019-01-31 23:59', '2 sec'::interval));
			INSERT INTO conn_log (user_id, ip_addr, ts) VALUES
				(1, '10.10.0.2', generate_series('2019-02-01 00:00'::timestamp, '2019-02-28 23:59', '2 sec'::interval));
			INSERT INTO conn_log (user_id, ip_addr, ts) VALUES
				(2, '10.10.0.3', generate_series('2019-03-01 00:00'::timestamp, '2019-03-31 23:59', '2 sec'::interval));
			INSERT INTO conn_log (user_id, ip_addr, ts) VALUES
				(3, '10.10.0.4', generate_series('2019-04-01 00:00'::timestamp, '2019-04-30 23:59', '2 sec'::interval));
			INSERT INTO conn_log (user_id, ip_addr, ts) VALUES
				(3, '10.10.0.5', generate_series('2019-05-01 00:00'::timestamp, '2019-05-31 23:59', '2 sec'::interval));
			INSERT INTO conn_log (user_id, ip_addr, ts) VALUES
				(4, '10.10.0.1', generate_series('2019-06-01 00:00'::timestamp, '2019-06-30 23:59', '2 sec'::interval));
			INSERT INTO conn_log (user_id, ip_addr, ts) VALUES
				(4, '10.10.0.3', generate_series('2019-07-01 00:00'::timestamp, '2019-07-31 23:59', '2 sec'::interval));
			INSERT INTO conn_log (user_id, ip_addr, ts) VALUES
				(5, '10.10.0.6', generate_series('2019-08-01 00:00'::timestamp, '2019-08-31 23:59', '2 sec'::interval));

      CREATE INDEX user_id_idx ON conn_log (user_id);

			CREATE TABLE user_id_to_ip AS
			SELECT DISTINCT user_id, ip_addr
			FROM conn_log;

			CREATE INDEX user_id_ip_idx ON user_id_to_ip (user_id);
		  `)
		if err != nil {
			return err
		}

		return err

	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			DROP TABLE if exists user_id_to_ip;
			DROP INDEX user_id_ip_idx;
			DROP INDEX user_id_idx;
			TRUNCATE TABLE conn_log;`)
		if err != nil {
			return err
		}
		return nil
	})
}
