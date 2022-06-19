package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Ulukbek-Toychuev/OpenWeather-Service/api"
	id "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd"
	s "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd"

	_ "github.com/lib/pq"
)

// Здесь будет код который работает с БД
// Here will be the code that works with the database
const (
	host     = "192.168.31.109"
	port     = 5432
	user     = "uluk"
	password = "rewq_1424"
	dbname   = "testdb"
)

func ConnectDB(city string) {

	currData, status := s.GetData(city)
	sonyaID := id.GetID()
	log.Println("Returned code status = ", status, "generated ID", sonyaID)

	psqInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected!")

	insertToTableLocation(db, &currData, &sonyaID)
	insertToTableCurrTemp(db, &currData, &sonyaID)
	insertToTableAirQuality(db, &currData, &sonyaID)
}

func insertToTableLocation(db *sql.DB, currData *api.Data, sonyaID *string) {
	sqlStatement := `
	INSERT INTO location (sonya_flakeid, city_name, country, lat, lon, local_time)
	VALUES (
		$1, $2, $3, $4, $5, $6
		); `

	_, err := db.Exec(sqlStatement, sonyaID, currData.Location.Name,
		currData.Location.Country, currData.Location.Lat, currData.Location.Lon,
		currData.Location.LocalTime)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully insert data to location table!")
	}
}

func insertToTableCurrTemp(db *sql.DB, currData *api.Data, sonyaID *string) {
	sqlStatement := `
	INSERT INTO current_temp (city_id, sonya_flakeid, temp_in_c, condition_text)
	VALUES (
		(SELECT location_id FROM location WHERE sonya_flakeid = $1), 
		(SELECT sonya_flakeid FROM location WHERE sonya_flakeid = $2),
		$3, $4); `

	_, err := db.Exec(sqlStatement, sonyaID, sonyaID, currData.Current.CurrentTemp,
		currData.Current.Condition.Text)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully insert data to current_temp table!")
	}
}

func insertToTableAirQuality(db *sql.DB, currData *api.Data, sonyaID *string) {
	sqlStatement := `
	INSERT INTO air_quality (city_id, sonya_flakeid, uv, co, no2, o3, so2, pm2_5, pm10, us_epa_index, gb_defra_index)
	VALUES(
		(select location_id from location where sonya_flakeid = $1), 
		(select sonya_flakeid from location where sonya_flakeid = $2),
		$3, $4, $5, $6, $7, $8, $9, $10, $11 
		); `

	_, err := db.Exec(sqlStatement, sonyaID, sonyaID, currData.Current.UV, currData.Current.AirQuality.Co,
		currData.Current.AirQuality.No2, currData.Current.AirQuality.O3, currData.Current.AirQuality.So2,
		currData.Current.AirQuality.Pm2_5, currData.Current.AirQuality.Pm10,
		currData.Current.AirQuality.US_epa_index, currData.Current.AirQuality.GB_defra_index)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully insert data to air_quality table!")
	}
}
