# Tasks in Golang
CRUD functionality in Golang and MySQL

## Getting Started
Clone this repository into your go home path or a separate directory.

Enter this command to install dependencies and run the project.
> glide install

After installing the dependencies, need to create `config.yml` file in the project root directory. It'll have our database credentials.
config.yml content,
```
database:
  dbname: "todos"
  dbuser: "root"
  dbpassword: "root123"
  dbhost: "127.0.0.1:3306"
```

We need to create a database to run our project, type below command in mysql client in order to create db,  
```CREATE DATABASE IF NOT EXISTS todos;```

After creating a database, we need to create a table which will store our data,
```
CREATE TABLE `todos`.`tasks` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(45) NOT NULL,
  `is_completed` TINYINT(1) NULL DEFAULT 0,
  `notify` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);
```

We are all set to run our project, type the below command
>go run main.go

