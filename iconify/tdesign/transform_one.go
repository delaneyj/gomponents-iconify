package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransformOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h6v2h8V2h6v6h-2v8h2v6h-6v-2H8v2H2v-6h2V8H2V2Zm4 6v8h2v2h8v-2h2V8h-2V6H8v2H6Zm0-4H4v2h2V4Zm14 2V4h-2v2h2Zm-2 12v2h2v-2h-2ZM6 20v-2H4v2h2Z"/>`),
		g.Group(children),
	)
}