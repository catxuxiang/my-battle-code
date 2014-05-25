package combat

import (
	"fmt"
	//	"time"
)

var attackerPos [5]point
var defenderPos [5]point

type point struct {
	x float64
	y float64
}

type BattleField struct {
	attacker *Team
	defender *Team
}

func (b *BattleField) SetAttacker(t *Team) {
	b.attacker = t
}

func (b *BattleField) SetDefender(t *Team) {
	b.defender = t
}

func (b *BattleField) Start() {
	attacker := &b.attacker.humans
	for i := 0; i < len(*attacker); i++ {
		(*attacker)[i].setPosition(attackerPos[i])
	}
	//fmt.Printf("attacker:%d\r\n", len(*attacker))
	defender := &b.defender.humans
	for i := 0; i < len(*defender); i++ {
		(*defender)[i].setPosition(defenderPos[i])
	}

	for i := 0; true; i++ {
		//time.Sleep(time.Second * 1)
		fmt.Printf("time:%d\r\n", i+1)
		if !b.attacker.loop(true, defender) || !b.defender.loop(false, attacker) {
			attackerDie := allDie(attacker)
			defenderDie := allDie(defender)
			if attackerDie && defenderDie {
				fmt.Printf("battle result:draw\r\n")
			} else if attackerDie {
				fmt.Printf("battle result:defender win\r\n")
			} else if defenderDie {
				fmt.Printf("battle result:attacker win\r\n")
			} else {
				fmt.Printf("battle result:time out\r\n")
			}
			break
		}
	}
}

func allDie(humans *[]*Human) bool {
	for i := 0; i < len(*humans); i++ {
		if (*humans)[i].Hp > 0 {
			return false
		}
	}
	return true
}

func init() {
	attackerPos[0] = point{50, 500}
	attackerPos[1] = point{30, 500}
	attackerPos[2] = point{30, 400}
	attackerPos[3] = point{10, 500}
	attackerPos[4] = point{10, 400}

	defenderPos[0] = point{100, 500}
	defenderPos[1] = point{890, 500}
	defenderPos[2] = point{890, 400}
	defenderPos[3] = point{870, 500}
	defenderPos[4] = point{870, 400}
}
