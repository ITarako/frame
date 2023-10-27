//go:build tools
// +build tools

package tools

import (
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
	_ "github.com/vektra/mockery/v2"
)

//go:generate go build -o ../bin/tools/mockery github.com/vektra/mockery/v2
//go:generate go build -tags 'postgres' -o ../bin/tools/migrate github.com/golang-migrate/migrate/v4/cmd/migrate
//go:generate chmod +x -R ../bin/tools
