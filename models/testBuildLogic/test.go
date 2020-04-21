package main

import (
	"com.waschild/noChaos-Server/models"
)

func main() {
	logic := models.NC_Logic{}
	logic.ID = 26
	models.NCDB.Debug().First(&logic)
	servlet := models.NC_Servlet{}
	servlet.ID = logic.ServletId
	models.NCDB.Debug().First(&servlet)
	logic.BuildLogic(servlet.GetName())
}
