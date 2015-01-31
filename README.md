# go-json-rest-postgresql-example
a simple example using Go to provide a json rest api with a postgresql database

this sample assumes you have a postgresql 9.x database setup.

I personally am using postgresql 9.4 with postgis.

I use https://github.com/ant0ine/go-json-rest for the rest interface

I use https://github.com/jackc/pgx for postgresql for its higher performance

I assume there is a schema called version1 and a table called locations.

The table has the following fields:

zipcode varchar 255
city varchar 255
state varchar 255
county varchar 255
country varchar 255

