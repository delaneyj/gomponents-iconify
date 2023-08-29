package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubscriptionsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V8h20v14H2Zm8-3l6-4l-6-4v8ZM4 7V5h16v2H4Zm3-3V2h10v2H7Z"/>`),
		g.Group(children),
	)
}