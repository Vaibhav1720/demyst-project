package todo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
	"fmt"
)
type item struct {
	UserId string
	Id string
	Title string
	Completed bool 
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		UserId: task,
		Id: time.Now().String(),
		Title :task,
		Completed: false,
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	ls[index-1].Completed = true

	return nil
}

func (t *Todos) Delete(index int) error  {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0{
		return err
	}

	err = json.Unmarshal(file, t)

	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func (t *Todos) PrintFirst20EnenNumberedTodos() {
	for index, item := range *t {
		index++
		if index%2 == 0 && index <=40 {
			if item.Completed {
				fmt.Printf("%d - %s -> Completed \n", index, item.Title)
			} else {
				fmt.Printf("%d - %s -> Pending \n", index, item.Title)
			}
		}
	}
}
