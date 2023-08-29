package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCardAlertOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 17h2v-2h-2v2Zm0-4h2V8h-2v5Zm9 9H4V8l6-6h10v20Zm-2-2V4h-7.15L6 8.85V20h12Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}