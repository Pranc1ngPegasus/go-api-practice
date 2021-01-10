// +build bin

package third_party

import (
	_ "github.com/cosmtrek/air"
	_ "github.com/google/wire"
	_ "github.com/rubenv/sql-migrate/sql-migrate"
	_ "github.com/volatiletech/sqlboiler"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql"
)

