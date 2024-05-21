package main

import (
	"database/sql"
	"log"

	p "github.com/Yandex-Practicum/go-db-sql-final/parcel"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open(p.DBDriver, p.PathDB)
	if err != nil {
		log.Println("can't open DB:", err)
	}
	defer db.Close()

	store := p.NewParcelStore(db)

	service := p.NewParcelService(store)

	// регистрация посылки
	client := 1
	address := "Псков, д. Пушкина, ул. Колотушкина, д. 5"
	p, err := service.Register(client, address)
	if err != nil {
		log.Println("Не удалось зарегистрировать посылку:", err)
		return
	}

	// изменение адреса
	newAddress := "Саратов, д. Верхние Зори, ул. Козлова, д. 25"
	err = service.ChangeAddress(p.Number, newAddress)
	if err != nil {
		log.Println("Не удалось изменить адрес посылки:", err)
		return
	}

	// изменение статуса
	err = service.NextStatus(p.Number)
	if err != nil {
		log.Println("Не удалось изменить статус посылки:", err)
		return
	}

	// вывод посылок клиента
	err = service.PrintClientParcels(client)
	if err != nil {
		log.Println("Не удалось поулчить список посылок:", err)
		return
	}

	// попытка удаления отправленной посылки
	err = service.Delete(p.Number)
	if err != nil {
		log.Println("Не удалось удалить посылку:", err)
		return
	}

	// вывод посылок клиента
	// предыдущая посылка не должна удалиться, т.к. её статус НЕ «зарегистрирована»
	err = service.PrintClientParcels(client)
	if err != nil {
		log.Println("Не удалось поулчить список посылок:", err)
		return
	}

	// регистрация новой посылки
	p, err = service.Register(client, address)
	if err != nil {
		log.Println("Не удалось зарегестрировать посылку:", err)
		return
	}

	// удаление новой посылки
	err = service.Delete(p.Number)
	if err != nil {
		log.Println("Не удалось удалить посылку:", err)
		return
	}

	// вывод посылок клиента
	// здесь не должно быть последней посылки, т.к. она должна была успешно удалиться
	err = service.PrintClientParcels(client)
	if err != nil {
		log.Println("Не удалось поулчить список посылок:", err)
		return
	}
}
