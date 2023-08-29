package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteSweepOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 18v-2h4v2h-4Zm0-8V8h7v2h-7Zm0 4v-2h6v2h-6ZM3 8H2V6h4V4.5h4V6h4v2h-1v9q0 .825-.588 1.413T11 19H5q-.825 0-1.413-.588T3 17V8Zm2 0v9h6V8H5Zm0 0v9v-9Z"/>`),
		g.Group(children),
	)
}