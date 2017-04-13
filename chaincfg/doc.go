// Package chaincfg defines chain configuration parameters.
//
// In addition to the main Decred network, which is intended for the transfer
// of monetary value, there also exists two currently active standard networks:
// regression test and testnet (version 0).  These networks are incompatible
// with each other (each sharing a different genesis block) and software should
// handle errors where input intended for one network is used on an application
// instance running on a different network.
//
// For library packages, chaincfg provides the ability to lookup chain
// parameters and encoding magics when passed a *Params.  Older APIs not updated
// to the new convention of passing a *Params may lookup the parameters for a
// wire.DecredNet using ParamsForNet, but be aware that this usage is
// deprecated and will be removed from chaincfg in the future.
//
// For main packages, a (typically global) var may be assigned the address of
// one of the standard Param vars for use as the application's "active" network.
// When a network parameter is needed, it may then be looked up through this
// variable (either directly, or hidden in a library call).
//
//  package main
//
//  import (
//          "flag"
//          "fmt"
//          "log"
//
//          "github.com/decred/dcrutil"
//          "github.com/looptrader/colonyd/chaincfg"
//  )
//
//  var testnet = flag.Bool("testnet", false, "operate on the testnet Decred network")
//
//  // By default (without -testnet), use mainnet.
//  var chainParams = &chaincfg.MainNetParams
//
//  func main() {
//          flag.Parse()
//
//          // Modify active network parameters if operating on testnet.
//          if *testnet {
//                  chainParams = &chaincfg.TestNetParams
//          }
//
//          // later...
//
//          // Create and print new payment address, specific to the active network.
//          pubKeyHash := make([]byte, 20)
//          addr, err := dcrutil.NewAddressPubKeyHash(pubKeyHash, chainParams)
//          if err != nil {
//                  log.Fatal(err)
//          }
//          fmt.Println(addr)
//  }
//
// If an application does not use one of the three standard Decred networks,
// a new Params struct may be created which defines the parameters for the
// non-standard network.  As a general rule of thumb, all network parameters
// should be unique to the network, but parameter collisions can still occur
// (unfortunately, this is the case with regtest and testnet sharing magics).
package chaincfg