package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCardAlertSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 17h2v-2h-2v2Zm0-4h2V8h-2v5Zm9 9H4V8l6-6h10v20Z"/>`),
		g.Group(children),
	)
}