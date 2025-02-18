package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("fulltostackSmartContract")

func main() {
	logger.SetLevel(shim.LogInfo)
	if err := shim.Start(new(SmartContract)); err != nil {
		logger.Error("Error starting fulltostackSmartContract - ", err)
	}
}
