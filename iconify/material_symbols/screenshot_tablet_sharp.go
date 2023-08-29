package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotTabletSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 20V4h22v16H1Zm5-2h12V6H6v12Zm7-1h4v-4h-1.5v2.5H13V17Zm-6-6h1.5V8.5H11V7H7v4Z"/>`),
		g.Group(children),
	)
}