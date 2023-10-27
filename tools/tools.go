//go:build tools
// +build tools

package tools

import (
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/vektra/mockery"
)

//go:generate go build -o ../bin/tools/mockery github.com/vektra/mockery/mockery
//go:generate go build -o ../bin/tools/migrate github.com/golang-migrate/migrate/v4/cmd/migrate
//go:generate chmod +x -R ../bin/tools

//go:generate echo Hello, Tools!
