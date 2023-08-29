package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 0 1-9 9m9-9a9 9 0 0 0-9-9m9 9H3m9 9a9 9 0 0 1-9-9m9 9c1.933 0 3.5-4.03 3.5-9S13.933 3 12 3m0 18c-1.933 0-3.5-4.03-3.5-9s1.567-9 3.5-9m-9 9a9 9 0 0 1 9-9"/>`),
		g.Group(children),
	)
}