package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h5m-5 0a9 9 0 0 0 9 9m-9-9a9 9 0 0 1 9-9m-4 9h8m-8 0c0 4.97 1.79 9 4 9m-4-9c0-4.97 1.79-9 4-9m4 9h5m-5 0c0-4.97-1.79-9-4-9m4 9c0 4.97-1.79 9-4 9m9-9a9 9 0 0 0-9-9m9 9a9 9 0 0 1-9 9"/>`),
		g.Group(children),
	)
}