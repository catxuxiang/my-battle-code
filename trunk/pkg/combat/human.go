package combat

import (
	"fmt"
)

const (
	StandBy     = "StandBy"
	Run         = "Run"
	AttackStart = "AttackStart"
	AttackValid = "AttackValid"
	AttackEnd   = "AttackEnd"
	Spell       = "Spell"
)

const (
	speed = 10
)

type IntervalAction struct {
	Name           string
	count          int
	Cd             int
	AttackRange    float64
	ChangeInterval []int
	ChangeState    []string
	ChangeAction   []func(count int, attacker *Human, defender *Human)
}

type Dot struct {
	human     *Human
	duration  int
	count     int
	cd        int
	dotAction func(attacker *Human, defender *Human)
}

type Human struct {
	team    *Team
	Name    string
	Hp      int
	Attack  int
	Defend  int
	State   string
	pos     point
	actions map[string]*IntervalAction
	Status  string
	dots    map[*Human]*Dot
}

func NewHuman() *Human {
	var i Human
	i.init()
	return &i
}

func (h *Human) AddDot(human *Human,
	duration int,
	cd int,
	dotAction func(attacker *Human, defender *Human)) {

	var dot Dot
	dot.human = human
	dot.duration = duration
	dot.cd = cd
	dot.dotAction = dotAction
	h.dots[human] = &dot
}

func (h *Human) init() {
	h.State = StandBy
	h.actions = make(map[string]*IntervalAction)
	h.dots = make(map[*Human]*Dot)
}

func (h *Human) GetAttackRange() float64 {
	return h.actions["Attack"].AttackRange
}

func (h *Human) showAction(s string) {
	fmt.Printf("pos:%f, %f\taction:%s(%d) %s\r\n", h.pos.x, h.pos.y, h.Name, h.Hp, s)
}

func (h *Human) runAction(action func(count int, attacker *Human, defender *Human), actionName string, enemy *Human, count int) bool {
	h.showAction(actionName + " " + enemy.Name)
	if action != nil {
		action(count, h, enemy)
		if h.Hp <= 0 || enemy.Hp <= 0 {
			return false
		}
	}
	return true
}

func (h *Human) loop(goRight bool, humans *([]*Human)) bool {
	for k, v := range h.dots {
		v.count++
		if v.count > v.duration {
			delete(h.dots, k)
		} else {
			if v.count%v.cd == 0 {
				v.dotAction(v.human, h)
			}
		}
	}

	dis := float64(9999)
	var enemy *Human
	for k, _ := range *humans {
		human := (*humans)[k]
		if human.Hp > 0 {
			distance := GetDistance(human.getPosition(), h.getPosition())
			//fmt.Printf("distance:%f\r\n", distance)
			if distance <= h.GetAttackRange() && distance < dis {
				dis = distance
				enemy = human
			}
		}
	}
	//如果能攻击到敌人
	if enemy != nil {
		if h.State == Run {
			h.State = StandBy
		}
		for k, _ := range h.actions {
			v := h.actions[k]
			v.count++
			if v.count > v.Cd {
				if v.count == v.Cd+1 && h.State != StandBy {
					v.count--
					h.showAction(h.State + "(change " + v.ChangeState[0] + " fail)")
				} else {
					tmp := v.Cd
					for k, _ := range v.ChangeInterval {
						tmp += v.ChangeInterval[k]
						if v.count <= tmp {
							h.State = v.ChangeState[k]
							if !h.runAction(v.ChangeAction[k], h.State, enemy, v.count-(tmp-v.ChangeInterval[k])) {
								return false
							}
							break
						}
					}
					if v.count > tmp {
						v.count = 0
						h.State = StandBy
						//h.showAction(h.State)
					}
				}
			}
		}
		if h.State == StandBy {
			h.showAction(h.State)
		}
	} else {
		if goRight {
			h.pos.x += speed
		} else {
			h.pos.x -= speed
		}
		if h.State == StandBy {
			h.State = Run
		}
		h.showAction(h.State)
		/*
			for k, _ := range h.actions {
				v := h.actions[k]
				v.count++
			}
		*/
	}
	return true
}

func (h *Human) AddAction(name string, cd int, attackRange float64, changeInterval *[]int, changeState *[]string, changeAction *[]func(count int, attacker *Human, defender *Human)) {
	var ia IntervalAction
	ia.Name = name
	ia.Cd = cd
	ia.AttackRange = attackRange
	ia.ChangeInterval = *changeInterval
	ia.ChangeState = *changeState
	ia.ChangeAction = *changeAction
	h.actions[ia.Name] = &ia
}

func (h *Human) setPosition(pos point) {
	h.pos.x = pos.x
	h.pos.y = pos.y
}

func (h *Human) getPosition() point {
	return h.pos
}
