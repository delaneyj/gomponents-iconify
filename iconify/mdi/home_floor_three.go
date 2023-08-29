package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeFloorThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 3l10 9h-3v8H5v-8H2l10-9m3 8.5V10a2 2 0 0 0-2-2H9v2h4v2h-2v2h2v2H9v2h4a2 2 0 0 0 2-2v-1.5a1.5 1.5 0 0 0-1.5-1.5a1.5 1.5 0 0 0 1.5-1.5Z"/>`),
		g.Group(children),
	)
}