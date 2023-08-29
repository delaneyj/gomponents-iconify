package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteBinFiveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 8h16v13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8Zm3-3V3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v2h5v2H2V5h5Zm2-1v1h6V4H9Zm0 8v6h2v-6H9Zm4 0v6h2v-6h-2Z"/>`),
		g.Group(children),
	)
}