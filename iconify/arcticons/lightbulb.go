package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightbulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.213 17.73a13.213 13.213 0 1 0-20.707 10.866a4.006 4.006 0 0 1 1.665 3.315v4.731h11.658v-4.735a4.051 4.051 0 0 1 1.703-3.337a13.174 13.174 0 0 0 5.681-10.84ZM19.444 38.928h9.112m-9.178 2.286h9.113m-6.97 2.286h4.958"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.011 36.642v-7.059l1.649-3.498m-5.671 10.557v-7.059l-1.649-3.498"/>`),
		g.Group(children),
	)
}