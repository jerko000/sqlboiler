package main

import (
	"github.com/jerko000/sqlboiler/drivers"
	"github.com/jerko000/sqlboiler/drivers/sqlboiler-mysql/driver"
)

func main() {
	drivers.DriverMain(&driver.MySQLDriver{})
}
