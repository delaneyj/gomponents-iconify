package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Midpoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 7.5v-3a2 2 0 0 0-2-2h-3m-3 10v-4m-2 2h4m6 3v3a2 2 0 0 1-2 2h-3m-6-16h-3a2 2 0 0 0-2 2v3m5 11h-3a2 2 0 0 1-2-2v-3"/>`),
		g.Group(children),
	)
}