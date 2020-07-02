/**
* Created by Maulana Ibrahim
* on 27 June 2020
 */

package main

import (
	api_master "maulibra/enigma_school/api-master"
	"maulibra/enigma_school/infrastructure"
)

func main() {
	//set environment via command line or setting manual default env in infrastructure config
	db := infrastructure.ConnectDB()
	router := infrastructure.NewRouter()
	api_master.Init(db, router)
	infrastructure.ListenServe(router)
}
