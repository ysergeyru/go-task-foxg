package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("Generating test data. Please wait...")
		_, err := db.Exec(`
      INSERT INTO conn_log (user_id, ip_addr, ts) values(
        generate_series(1,4000000)::bigint,
        host(set_masklen(((generate_series(1, (2 ^ (32 - masklen('10.10.0.0/10'::inet)))::integer - 2) + '10.10.0.0/10'::inet)::inet), 32)),
        generate_series('2017-01-01 00:00'::timestamp, '2017-12-31 23:50', '8 sec'::interval));
      INSERT INTO conn_log (user_id, ip_addr, ts) values(
        generate_series(2,4000001)::bigint,
        host(set_masklen(((generate_series(1, (2 ^ (32 - masklen('10.10.0.0/10'::inet)))::integer - 2) + '10.10.0.0/10'::inet)::inet), 32)),
        generate_series('2018-01-01 00:00'::timestamp, '2018-12-31 23:50', '8 sec'::interval));
      INSERT INTO conn_log (user_id, ip_addr, ts) values(
        generate_series(3,4000002)::bigint,
        host(set_masklen(((generate_series(1, (2 ^ (32 - masklen('10.10.0.0/10'::inet)))::integer - 2) + '10.10.0.0/10'::inet)::inet), 32)),
        generate_series('2019-01-01 00:00'::timestamp, '2019-12-31 23:50', '8 sec'::interval));

      CREATE INDEX user_id_idx ON conn_log (user_id);
      `)
		return err
	}, func(db migrations.DB) error {
		return nil
	})
}
