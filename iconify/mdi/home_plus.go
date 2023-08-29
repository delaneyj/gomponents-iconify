package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 14h2v3h3v2h-3v3h-2v-3h-3v-2h3v-3M12 3l10 9h-4a6.005 6.005 0 0 0-5.66 8H5v-8H2l10-9Z"/>`),
		g.Group(children),
	)
}