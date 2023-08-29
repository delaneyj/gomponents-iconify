package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCopySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19V1h9l6 6v12H6Zm8-11h5l-5-5v5ZM2 23V7h2v14h11v2H2Z"/>`),
		g.Group(children),
	)
}