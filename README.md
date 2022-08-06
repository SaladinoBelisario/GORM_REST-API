# Basic REST API with GORM and PostgreSQL connection

This is a basic REST API made in Golang using GORM for 
automatic table management and connected to a PostgreSQL Docker
container.

The Postgre default configuration is provided in [Docker compose](docker-compose.yaml),
but you can also provide a .env file and use it with the project as already contains the 
dependencies to use .env files.

The [Air TOML](.air.toml) file is for using Golang Air package to avoid redeploy manually
every time you need to change the code in development phase.