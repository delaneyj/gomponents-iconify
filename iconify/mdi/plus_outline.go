package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 9h5V4h6v5h5v6h-5v5H9v-5H4V9m7 4v5h2v-5h5v-2h-5V6h-2v5H6v2h5Z"/>`),
		g.Group(children),
	)
}