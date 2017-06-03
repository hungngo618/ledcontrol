package db

import (
	log "github.com/Sirupsen/logrus"
)

func GetLedStatus() (int, error) {
	var id int
	err := DB.QueryRow("SELECT led_on FROM `ledcontrol` WHERE id=1;").Scan(&id)
	if err != nil {
		log.Error("failed to count intents in database: ", err)
		return 0, err
	}

	return id, nil
}

func TurnLedOn()( error) {
	query := "UPDATE ledcontrol SET led_on=1 where id=1;"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare query: ", err)
		return err
	}
	defer stmt.Close()

	// TODO: Use transaction
	_, err = stmt.Exec()
	if err != nil {
		log.Error("failed to execute query: ", err)
		return err
	}
	return nil
}

func TurnLedOff()( error) {
	query := "UPDATE ledcontrol SET led_on=0 where id=1;"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare query: ", err)
		return err
	}
	defer stmt.Close()

	// TODO: Use transaction
	_, err = stmt.Exec()
	if err != nil {
		log.Error("failed to execute query: ", err)
		return err
	}
	return nil
}
