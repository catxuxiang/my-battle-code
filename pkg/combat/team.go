package combat

import (
//	"fmt"
)

type Team struct {
	humans []*Human
}

func (t *Team) AddHuman(h *Human) {
	t.humans = append(t.humans, h)
	h.team = t
}

func (t *Team) loop(goRight bool, humans *([]*Human)) bool {
	for _, v := range t.humans {
		if v.Hp > 0 {
			//fmt.Printf("%s\r\n", v.Name)
			if !v.loop(goRight, humans) {
				return false
			}
		}
	}
	return true
}
