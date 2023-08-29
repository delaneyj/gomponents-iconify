package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m369 81l30 30l-30 31q42 52 42 119q0 58-32 106l-31-31q20-35 20-75q0-62-43.5-105.5T219 112q-40 0-75 20l-31-31q48-32 106-32q67 0 120 42zM283 5v43H155V5h128zm-86 180v-30h43v73zM27 69l214 214l164 165l-27 27l-53-54q-48 32-106 32q-80 0-136-56T27 261q0-58 32-106L0 96zm192 342q40 0 75-21L90 186q-21 35-21 75q0 62 44 106t106 44z"/>`),
		g.Group(children),
	)
}