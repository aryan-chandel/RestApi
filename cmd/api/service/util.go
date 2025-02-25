package service

import (
	"encoding/json"
	"os"

	"github.com/labstack/echo/v4"
)

type Dta struct {
	UserId int
	Id     int
	Title  string
	Body   string
}
type payload struct {
	Data []Dta //slice of dta struct therefore interface implements list of data node

}

func raw() ([]Dta, error) {
	r, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	var pload payload
	err = json.Unmarshal(r, &pload)
	if err != nil {
		return nil, err
	}
	return pload.Data, nil

}
func GetAll() ([]Dta, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}
	return data, nil

}
func GetById(index int) (any, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}
	if index > len(data) {
		res := make([]string, 0) // sending empty string
		return res, nil
	}
	return data[index-1], nil
}
func Postuser(c echo.Context) error {
	var newel Dta
	err := json.NewDecoder(c.Request().Body).Decode(&newel)
	if err != nil {
		return err
	}
	data, err := raw()
	if err != nil {
		return err
	}
	data = append(data, newel)
	Data, err := json.Marshal(data)
	if err != nil {
		return err
	}

	er := os.WriteFile("data.json", Data, 0644)
	if er != nil {
		return er
	}
	return nil
}
func DeleteUser(id int) error {
	data, err := GetAll()
	if err != nil {
		return err
	}
	var indx = -1
	for idd, user := range data {
		if user.Id == id {
			indx = idd
			break
		}
	}
	if indx == -1 {
		var err error

		return err
	}
	data = append(data[:indx], data[indx+1:]...)
	body, er := json.MarshalIndent(data, "", " ")
	if er != nil {
		return er
	}
	pr := os.WriteFile("data.json", body, 0644)
	if pr != nil {
		return pr
	}
	return nil
}
