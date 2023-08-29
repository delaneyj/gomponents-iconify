package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Finance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V3h2v16h16v2H5Zm1-3V9h4v9H6Zm5 0V4h4v14h-4Zm5 0v-5h4v5h-4Z"/>`),
		g.Group(children),
	)
}