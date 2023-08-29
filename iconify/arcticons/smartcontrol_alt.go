package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartcontrolAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.5 8.5c-15.5 0-15.5 31-31 31"/>`),
		g.Group(children),
	)
}