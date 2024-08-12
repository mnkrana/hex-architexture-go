package main

import (
	"log"
	"os"

	"github.com/mnkrana/hex/internal/adapters/app/api"
	"github.com/mnkrana/hex/internal/adapters/core/arithmetic"
	"github.com/mnkrana/hex/internal/adapters/framework/left/grpc"
	"github.com/mnkrana/hex/internal/adapters/framework/right/db"
	"github.com/mnkrana/hex/internal/ports"
)

// "fmt"

func main() {

	//ports
	var dbAdapter ports.DbPort
	var core ports.ArithmaticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	//core port
	core = arithmetic.New()

	//db port
	dbDriver := os.Getenv("DB_DRIVER")
	dbSourceName := os.Getenv("DB_SOURCE_NAME")
	dbAdapter, err := db.NewAdapter(dbDriver, dbSourceName)
	if err != nil {
		log.Fatalf("failed to initiate db connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	//app port
	appAdapter = api.NewAdaptor(dbAdapter, core)

	//grpc port
	gRPCAdapter = grpc.NewAdaptor(appAdapter)
	gRPCAdapter.Run()
}
