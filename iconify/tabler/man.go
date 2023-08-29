package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Man(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 16v5m4-5v5M9 9h6l-1 7h-4zm-4 2c1.333-1.333 2.667-2 4-2m10 2c-1.333-1.333-2.667-2-4-2m-5-5a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/>`),
		g.Group(children),
	)
}