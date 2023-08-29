package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256.286 408.357L16.333 138.548V104H496v36.45ZM56.892 136l199.466 224.287L457.042 136Z"/>`),
		g.Group(children),
	)
}