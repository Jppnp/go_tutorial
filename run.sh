#!/bin/sh
dbUrl="jdbc:postgresql://postgres:5432/postgres"
dbUserName="postgres"
dbPassword=""
defaultSchemaName="go_tutorial"

cd migrations && java -jar liquibase.jar \
	--driver=org.postgresql.Driver \
	--classpath=./lib/postgresql.jar \
	--url=$dbUrl \
	--changeLogFile=./changelog/changelog-initschema.xml \
	--username=$dbUserName\
	--password=$dbPassword\
	--defaultSchemaName=public\
	--databaseChangeLogLockTableName="${PROJECT_NAME}_databasechangeloglock"\
	--databaseChangeLogTableName="${PROJECT_NAME}_databasechangelog"\
	update;
java -jar liquibase.jar \
	--driver=org.postgresql.Driver \
	--classpath=./lib/postgresql.jar \
	--url=$dbUrl \
	--changeLogFile=./changelog/changelog-master.xml \
	--username=$dbUserName\
	--password=$dbPassword\
	--defaultSchemaName=$defaultSchemaName\
	update;

echo "Starting server..."
echo "Project : ${PROJECT_NAME}"
echo "Environment : ${BUILD_STAGE}"
cd ..
# Run the built executable
./${PROJECT_NAME}

