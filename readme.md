# GO-ASSESMENT
Multiple choice based assesment with [golang](https://golang.org/)

## Techstack
1. [gin](https://gin-gonic.com/)
2. [gorm](https://gorm.io/)
3. [MySQL](https://www.mysql.com)

## How to install 
1. You have to installed golang 1.17 or greater version and MySQL.
2. Create database and configure the database connection at `/helper/db.go`.
3. Install the required modules by run `go mod install`.
4. Run the project by run `go run main.go`.
5. Go to `http://127.0.0.1:8080`.

## Project Directory Structure
* [controller](./controller)
* [helper](./helper)
* [message](./message)
* [migration](./migration)
* [model](./model)
* [route](./route)
* [go.mod](./go.mod)
* [go.sum](./go.sum)
* [main.go](./main.go)
* [readme.md](./readme.md)

### controller
The directory where all handler live.

### helper
Contains modules which can be reuseable in another modules.

### message
Contains the constant value for API message response.

### migrations
Contains the migrations which built by `gorm`.

### model
The directory where models stored.

### route
The diretory where the project route registered.

### go.mod
A file which contain installed modules.

### go.sum
A file which lists down the checksum of direct and indirect dependency required by the module.

### main.go
The main file where the project executed.

### readme.md
A file for documentation purpose.