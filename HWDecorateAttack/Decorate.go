package main

import (
	"fmt"
	"math/rand/v2"
)

func Attack() string {
	return "Атака выполнена!"
}

func DamageBoostDecorator(attackFunc func() string) func() string {
	return func() string {
		return fmt.Sprintf("Вам улыбнулась удача, нанесение урона увеличено на 10%%! ") + attackFunc()
	}
}

func CriticalHitDecorator(attackFunc func() string) func() string {
	return func() string {
		if rand.IntN(4) == 3 {
			return "Критический удар! Урон удвоен! " + attackFunc()
		}
		return attackFunc()
	}
}

func SlowEffectDecorator(attackFunc func() string) func() string {
	return func() string {
		return attackFunc() + " Цель замедлена на 2 хода!"
	}
}

func main() {
	attackWithDamageBoostDecorator := DamageBoostDecorator(Attack)
	attackWithCriticalHitDecorator := CriticalHitDecorator(attackWithDamageBoostDecorator)
	getSlowEffectDecorator := SlowEffectDecorator(attackWithCriticalHitDecorator)
	fmt.Println(getSlowEffectDecorator())
}
