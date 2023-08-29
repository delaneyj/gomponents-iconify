package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LampWallDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 13h6l3 7H8l3-7Zm3 0V8a2 2 0 0 0-2-2H8M4 9h2a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H4v6Z"/>`),
		g.Group(children),
	)
}