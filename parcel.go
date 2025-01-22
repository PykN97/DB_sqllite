package main

import (
	"database/sql"
	"fmt"
)

type ParcelStore struct {
	db *sql.DB
}

func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

func (s ParcelStore) Add(p Parcel) (int, error) {
	res, err := s.db.Exec("INSERT INTO parcel (client,status,address,created_at) VALUES (?,?,?,?)",
		p.Client, p.Status, p.Address, p.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// верните идентификатор последней добавленной записи
func (s ParcelStore) Get(number int) (Parcel, error) {
	row := s.db.QueryRow("SELECT number,client,status,address,created_at FROM parcel WHERE number = ?", number)
	// реализуйте чтение строки по заданному number
	// здесь из таблицы должна вернуться только одна строка

	// заполните объект Parcel данными из таблицы
	p := Parcel{}
	err := row.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
	if err != nil {
		return Parcel{}, err
	}
	return p, err
}

func (s ParcelStore) GetByClient(client int) ([]Parcel, error) {
	rows, err := s.db.Query("SELECT number,client,status,address,created_at FROM parcel WHERE client = ?", client)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	var res []Parcel
	for rows.Next() {
		var p Parcel
		err := rows.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

// реализуйте чтение строк из таблицы parcel по заданному client
// здесь из таблицы может вернуться несколько строк

// заполните срез Parcel данными из таблицы

func (s ParcelStore) SetStatus(number int, status string) error {
	_, err := s.db.Exec("UPDATE parcel SET status = ? WHERE number = ?", status, number)
	if err != nil {
		fmt.Println(err)
		return err // Возвращаем ошибку
	}
	return nil
}
func (s ParcelStore) SetAddress(number int, address string) error {
	_, err := s.db.Exec("UPDATE parcel SET address = ? WHERE number = ? AND status = ?", address, number, ParcelStatusRegistered)
	if err != nil {
		return err
	}
	return nil
}

// реализуйте обновление адреса в таблице parcel
// менять адрес можно только если значение статуса registered

func (s ParcelStore) Delete(number int) error {
	_, err := s.db.Exec("DELETE FROM parcel WHERE status = ? AND number = ?", ParcelStatusRegistered, number)
	if err != nil {
		return err
	}
	return nil
}

// реализуйте удаление строки из таблицы parcel
// удалять строку можно только если значение статуса registered
