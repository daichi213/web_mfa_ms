#!/bin/sh

echo "Starting DB creation for testing and production..."

messages=`/usr/local/bin/psql /docker-entrypoint-initdb.d/create_test_db.sql`

res=$?

if [ $res -eq 0 ]; then
    echo "DB created successfully"
else
    echo "DB creation failed"
    echo $messages
fi
