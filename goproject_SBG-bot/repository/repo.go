package repository

import (
	"errors"
	"fmt"
	"goproject_SBG-bot/data"
	"goproject_SBG-bot/datastruct"
	"strconv"
	"time"
)

const (
	file = "tmp/Birthbay.csv"
)

type read interface {
	ReadFile_2() ([][]string, error)
}

type Repository struct {
	Persons_id   map[int64]*datastruct.Person
	Persons_name map[string]*datastruct.Person
	cnt          int
}

func New() *Repository {
	return (&Repository{}).Init()
}

func (r *Repository) Init() *Repository {
	reader := data.NewReader(file)
	record, err := ReadFile(reader)
	if err != nil {
		fmt.Println(err) // может return nil, err
	}

	map_person_n := map[string]*datastruct.Person{}

	for _, v := range record {
		if len(v) > 1 {
			p := &datastruct.Person{}
			p.Name = v[0]
			//			i, err := strconv.ParseInt(s, 10, 64)
			p.ID, _ = strconv.ParseInt(v[1], 10, 64)
			p.Date, _, _ = checkDate(v[2])
			//			p.Month, _ = strconv.Atoi(v[3])
			//			p.Day, _ = strconv.Atoi(v[4])

			m := map[string]int{}
			for i := 4; i < len(v); i++ {
				m[v[i]] = 1
			}
			p.Subscribers = m
			map_person_n[p.Name] = p
		}
	}
	//++++++++++++++++++++++++++++++++++++++++++++++++++

	map_person_id := map[int64]*datastruct.Person{}
	for _, v := range map_person_n {
		if v.ID != 0 {
			map_person_id[v.ID] = v
			fmt.Println(*v)
		}
	}
	//	fmt.Println(map_person_n["Gdgrd"])
	r.Persons_id = map_person_id
	r.Persons_name = map_person_n

	return r
}

func (r *Repository) Chek_avtorisation(chatID int64) bool {
	_, ok := r.Persons_id[chatID]
	if ok == false {
		p := &datastruct.Person{}
		p.ID = chatID
		r.Persons_id[chatID] = p
	}

	if r.Persons_id[chatID].Date == "" {
		return false
	} else {
		return true
	}
}

func (r *Repository) EnterName(text_vvod string, chatID int64) error {
	p := r.Persons_id[chatID]
	if p.Name == "" && p.Previous != "name" {
		p.Previous = "name"
		return errors.New("Авторизация введите имя")
	}

	if p.Previous == "name" {

		if checkName(text_vvod) == false {
			return errors.New("Авторизация введите имя")
		}
		p.Name = text_vvod
		p.Previous = ""
		return nil
	}
	return nil
}

func (r *Repository) EnterDate(text_vvod string, chatID int64) error {
	p := r.Persons_id[chatID]
	if p.Date == "" && p.Previous != "year" {
		p.Previous = "year"
		return errors.New("Авторизация введите дату рождения")
	}

	if p.Previous == "year" {
		date, _, err := checkDate(text_vvod)
		if err != nil {
			return errors.New("Авторизация введите дату рождения")
		}
		p.Date = date
		p.Previous = ""
		r.Persons_name[p.Name] = p
		p.Subscribers = map[string]int{}
		return nil
	}
	return nil
}

func (r *Repository) Out_list(chatID int64) string {

	textout := "Зарегестрированы следующие сотрудники:"
	for n := range r.Persons_name {
		textout = textout + "\n" + n
	}

	return textout
}

func (r *Repository) Get(chatID string) datastruct.Person {
	return *r.Persons_name[chatID]
}

func checkName(s string) bool {

	for _, v := range s {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) || (v >= 1040 && v <= 1071) || (v >= 1072 && v <= 1103) {
		} else {
			return false
		}
	}
	return true
}

func checkDate(text_vvod string) (string, time.Time, error) {
	var t time.Time
	var err error

	const shortForm = "2006-01-02"
	t, err = time.Parse(shortForm, text_vvod)
	if err != nil {
		return "", t, err
	}
	//	y := t.Date()
	return text_vvod, t, nil
}

func (r *Repository) Get_previous(chatID int64) string {
	return r.Persons_id[chatID].Previous
}

func (r *Repository) AddName(chatID int64) error {

	_, ok := r.Persons_id[chatID]
	if ok == true {
		r.Persons_id[chatID].Previous = "Add"
		return nil

	} else {
		return errors.New("Error")
	}

	return nil

}

func (r *Repository) AddNameWork(text_vvod string, chatID int64) error {

	_, q := r.Persons_name[text_vvod]
	if q == true {
		r.Persons_id[chatID].Subscribers[text_vvod] = 1
		r.Persons_id[chatID].Previous = ""
		return nil
	} else {
		return errors.New("Error")
	}

	return nil
}

func (r *Repository) DeleteName(chatID int64) (string, error) {
	var listName string

	p := r.Persons_id[chatID]
	//	n := 0
	for p_name := range p.Subscribers {
		//		n++
		listName = listName + "\n" + p_name
	}

	if listName == "" {
		return listName, errors.New("Error")
	}
	p.Previous = "Delete"

	return listName, nil
}

func (r *Repository) DeleteNameWork(text_vvod string, chatID int64) error {

	p := r.Persons_id[chatID]
	_, q := r.Persons_name[text_vvod]
	if q == false {
		return errors.New("Error")
	} else {
		_, ok := p.Subscribers[text_vvod]
		if ok == true {
			delete(p.Subscribers, text_vvod)
			p.Previous = ""
			return nil

		} else {
			return errors.New("Error")
		}
	}

	return nil
}

func (r *Repository) Сancel(chatID int64) string {
	var textout string
	_, ok := r.Persons_id[chatID]
	if ok == true {
		r.Persons_id[chatID].Previous = ""
		textout = "Отменено"
	}
	return textout
}

func (r *Repository) Get_worker() ([][]int64, []string) {

	to_name := make([]string, 0, 10)
	slice_ID := make([][]int64, 0, 10)
	t_now := time.Now()
	year := t_now.Year()
	for _, p := range r.Persons_id {
		if p.Date != "" {
			_, t_birthday, _ := checkDate(p.Date)
			day := t_birthday.Day()
			month := t_birthday.Month()
			//			fmt.Println("name ", p.name)
			t_person := time.Date(year, time.Month(month), day, 23, 59, 0, 0, time.UTC)
			if t_person.After(t_now) == false {
				t_person = t_person.AddDate(1, 0, 0)
			}
			t_hour := t_person.Sub(t_now).Hours()
			if t_hour < 24 {
				to_name = append(to_name, p.Name)
				slice_ID = append(slice_ID, r.find_list(p.Name))
			}
		}
	}

	return slice_ID, to_name
}

func (r *Repository) find_list(name string) []int64 {

	list := make([]int64, 0, 10)

	for _, p := range r.Persons_id {
		if p.Subscribers[name] == 1 {
			list = append(list, p.ID)
		}
	}
	return list
}

func ReadFile(reader read) ([][]string, error) {

	return reader.ReadFile_2()
}
