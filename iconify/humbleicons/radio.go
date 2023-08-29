package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.636 18.364a9 9 0 0 1 0-12.728m12.728 0a9 9 0 0 1 0 12.728M8.111 15.889a5.5 5.5 0 0 1 0-7.778m7.778 0a5.5 5.5 0 0 1 0 7.778M14 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/>`),
		g.Group(children),
	)
}