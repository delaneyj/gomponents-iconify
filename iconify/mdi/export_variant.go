package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExportVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1L8 5h3v9h2V5h3m2 18H6a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h3v2H6v12h12V9h-3V7h3a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}