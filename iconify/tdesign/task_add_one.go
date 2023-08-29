package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskAddOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0 0 18h1v2h-1C5.925 23 1 18.075 1 12S5.925 1 12 1c1.498 0 2.928.3 4.232.844l.923.385l-.77 1.846l-.923-.385A8.97 8.97 0 0 0 12 3Zm11.414 1.5L12 15.914L6.086 10L7.5 8.586l4.5 4.5l10-10L23.414 4.5ZM20 15v3h3v2h-3v3h-2v-3h-3v-2h3v-3h2Z"/>`),
		g.Group(children),
	)
}