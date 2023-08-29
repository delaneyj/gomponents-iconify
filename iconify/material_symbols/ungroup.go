package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ungroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.7 16.7l-1.4-1.4l4.3-4.3H8V9h7v7h-2v-3.6l-4.3 4.3ZM19 12V5h-7V3h9v9h-2ZM5 21q-.825 0-1.413-.588T3 19V3h2v16h16v2H5Z"/>`),
		g.Group(children),
	)
}