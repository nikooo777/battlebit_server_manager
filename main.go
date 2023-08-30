package main

import (
	"battlebit_admin_panel/cmd"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate go-bindata -o migration/bindata.go -nometadata -pkg migration -ignore bindata.go migration/
//go:generate go fmt ./migration/bindata.go
//go:generate goimports -l ./migration/bindata.go
func main() {
	cmd.Execute()
}
