package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightBoldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 16H3V8h8V2l10 10l-10 10v-6m2-9v3H5v4h8v3l5-5l-5-5Z"/>`),
		g.Group(children),
	)
}