package main

import (
	"github.com/jerko000/sqlboiler/drivers"
	"github.com/jerko000/sqlboiler/drivers/sqlboiler-psql/driver"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
