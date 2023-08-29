package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRollOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 5v15h-8v2H2V3h3V1h6v2h3v2h8ZM4 20h8v-2h8V7h-8V5H4v15Zm5-3h2v-2H9v2Zm0-7h2V8H9v2Zm4 7h2v-2h-2v2Zm0-7h2V8h-2v2Zm4 7h2v-2h-2v2Zm0-7h2V8h-2v2ZM4 20V5v15Z"/>`),
		g.Group(children),
	)
}