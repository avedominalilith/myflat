/*
Описать дом
1. Величины
2. Мебель 5 объектов
3. Бытовая техника 5 объектов
4. Состав семьи
5. Одна функция вызова, которой можно увидеть все
6. Выложить в гитхаб
*/
package main

import (
	"fmt"
	"strings"
)

type Flat struct {
	floor      int64
	area       float64
	numRooms   int64
	cost       int64
	whatInside whatInside
	family     family
}
type whatInside struct {
	appliance []string
	devices   []string
}
type family struct {
	members []string
	pets    []string
}

func main() {
	flat1 := Flat{
		floor:    1,
		area:     54,
		numRooms: 2,
		cost:     999999999,
		whatInside: whatInside{
			appliance: []string{"Диван", "Кровать", "Турник", "Шкаф", "Стул"},
			devices:   []string{"Компьютер,", "Ноут,", "Телевизор,", "Умный пылесос,", "Плэйстейшен"},
		},
		family: family{
			members: []string{"Мама,", "Папа,", "Брат,", "Кот,", "Собака"},
			pets:    []string{"Кузя,", "Сиша"},
		},
	}
	flat1.showFlat()
}
func (h Flat) showFlat() {
	fmt.Println("Я живу на", h.floor, "этаже")
	fmt.Println("Площадь моей квартиры около", h.area, "кв.см")
	fmt.Println("Всего комнат = ", h.numRooms)
	fmt.Println("Цена = ", h.cost, "рублей")
	fmt.Println("Моя семья:", strings.Trim(fmt.Sprint(h.family.members), "[]"))
	fmt.Println("Мои питомцы:", strings.Trim(fmt.Sprint(h.family.pets), "[]"))
	fmt.Println("Моя техника:", strings.Trim(fmt.Sprint(h.whatInside.appliance), "[]"))
	fmt.Println("Моя мебель:", strings.Trim(fmt.Sprint(h.whatInside.devices), "[]"))

}
