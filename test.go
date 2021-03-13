package main

import (
	"AvitoTask/internal/app/statistic/repository"
	"AvitoTask/internal/pkg/utility"
	"fmt"
)

func main() {
	sqlCon, err := utility.CreatePostgresConnection("host=localhost port=5432 dbname=avitoadvertising sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	statRep := repository.NewStatisticRepository(sqlCon)

	//stat := models.Statistic{
	//	Date:   "2010-07-13",
	//	Views:  50,
	//	Clicks: 200,
	//	Cost:   decimal.NewFromFloat(1000.30),
	//}
	//err = statRep.SaveStatistic(stat)
	//err = statRep.DeleteAll()
	stats, err := statRep.GetStatistic("2006-07-13", "2009-07-13", "cpm")
	for _, item := range stats {
		fmt.Println(item)
	}
	if err != nil {
		fmt.Println(err)
	}
}
