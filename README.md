# REST service

It accepts two user ids and respond with boolean value, which will be "true" if these users are duplicates.

Duplicates are pairs of user_ids for which there are at least two matching ip addresses in the access log.

## Install
Clone repository to your local machine.

Run instance of PostgreSQL and create db named 'service_data'.

Fill credentials for db user in file ./config/development/config.yaml

Run 'prepare_db_test_data.sh' to generate test data.

Be patient, generating of more than 10 million records in 'conn_log' table can take some time.

Now you can run project by ./rundev.sh script

And benchmark test with ./run_test_bench.sh

Postman tests are here ./postman/
