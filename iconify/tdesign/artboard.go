package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Artboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2v4h8V2h2v4h4v2h-4v8h4v2h-4v4h-2v-4H8v4H6v-4H2v-2h4V8H2V6h4V2h2Zm0 6v8h8V8H8Z"/>`),
		g.Group(children),
	)
}