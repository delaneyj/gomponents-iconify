package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReorderVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 3v18h2V3H9M5 3v18h2V3H5m8 0v18h2V3h-2m6 0h-2v18h2V3Z"/>`),
		g.Group(children),
	)
}