package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m54 364l39-39l30 30l-39 39zm159 83v-63h43v63h-43zM64 192v43H0v-43h64zm235-89q29 17 46.5 46t17.5 64q0 53-37.5 90.5T235 341t-90.5-37.5T107 213q0-35 17-64t47-46V0h128v103zm106 89h64v43h-64v-43zm-59 163l30-29l39 38l-30 30z"/>`),
		g.Group(children),
	)
}