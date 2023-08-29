package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 8.5a3.5 3.5 0 1 0 3.5 3.5A3.5 3.5 0 0 0 16 8.5Zm0 5a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5ZM16 5H8a7 7 0 0 0 0 14h8a7 7 0 0 0 0-14Zm0 12H8A5 5 0 0 1 8 7h8a5 5 0 0 1 0 10Z"/>`),
		g.Group(children),
	)
}