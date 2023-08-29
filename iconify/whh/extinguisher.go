package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Extinguisher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m814 60l-242 99q-5 18-14 33h242q13 0 22.5 9.5T832 224t-9.5 22.5T800 256H512v72q84 22 138 91t54 157v384q0 27-19 45.5t-45 18.5H256q-27 0-45.5-18.5T192 960V576q0-88 54-157t138-91v-72H256q-53 0-90.5 37.5T128 384v512q0 27-18.5 45.5t-45 18.5T19 941.5T0 896V384q0-106 75-181t181-75h64q0-53 37.5-90.5T448 0q42 0 75.5 25T570 90L778 4q13-6 27.5-2.5T827 16t3.5 24T814 60z"/>`),
		g.Group(children),
	)
}