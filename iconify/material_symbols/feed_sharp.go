package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h13l5 5v13H3ZM15 5v4h4l-4-4ZM7 17h10v-2H7v2Zm0-8h5V7H7v2Zm0 4h10v-2H7v2Z"/>`),
		g.Group(children),
	)
}