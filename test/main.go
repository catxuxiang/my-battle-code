// main
package main

import (
	"fmt"

	"pkg/combat"
)

func main() {
	var attacker combat.Team
	dz := combat.NewHuman()
	dz.Name = "dz"
	dz.Hp = 10
	dz.Attack = 10
	dz.Defend = 5
	dz.State = combat.StandBy
	dz.AddAction("Attack",
		5,
		10,
		&([]int{3, 2, 1}),
		&([]string{combat.AttackStart, combat.AttackValid, combat.AttackEnd}),
		&([]func(count int, attacker *combat.Human, defender *combat.Human){
			nil,
			func(count int, attacker *combat.Human, defender *combat.Human) {
				fmt.Printf("%s action(%d) %s", attacker.Name, count, defender.Name)
				if count == 2 {
					defender.Hp -= attacker.Attack - defender.Defend
					fmt.Printf(":%d hurts\r\n", attacker.Attack-defender.Defend)
				} else {
					fmt.Printf("\r\n")
				}
			},
			nil}))
	dz.AddAction("Spell",
		7,
		20,
		&([]int{2}),
		&([]string{combat.Spell}),
		&([]func(count int, attacker *combat.Human, defender *combat.Human){
			func(count int, attacker *combat.Human, defender *combat.Human) {
				if count == 2 {
					fmt.Printf("%s give a dot to %s\r\n", attacker.Name, defender.Name)
					defender.AddDot(attacker,
						10,
						2,
						func(attacker *combat.Human, defender *combat.Human) {
							defender.Hp -= 2
							fmt.Printf("%s suffer %d dot hurts by %s\r\n", defender.Name, -2, attacker.Name)
						})
				}
			},
		}))
	attacker.AddHuman(dz)
	var defender combat.Team
	dz1 := combat.NewHuman()
	dz1.Name = "ss"
	dz1.Hp = 20
	dz1.Attack = 15
	dz1.Defend = 2
	dz1.State = combat.StandBy
	dz1.AddAction("Attack", 5, 10, &([]int{3}), &([]string{combat.AttackValid}), &([]func(count int, attacker *combat.Human, defender *combat.Human){nil}))
	dz1.AddAction("Spell", 7, 20, &([]int{2}), &([]string{combat.Spell}), &([]func(count int, attacker *combat.Human, defender *combat.Human){nil}))
	defender.AddHuman(dz1)

	var field combat.BattleField
	field.SetAttacker(&attacker)
	field.SetDefender(&defender)
	field.Start()
}
