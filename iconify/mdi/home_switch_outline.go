package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSwitchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 3L1 9h2v6h4v-4h2v4h4V9h2L8 3m3.5 6v4.5h-1v-4h-5v4h-1V8L8 5l3.5 3v1M9 16v2h6v-2l3 3l-3 3v-2H9v2l-3-3l3-3m14-7h-2v6h-6v-5h4l-5.46-4.89L16 3l7 6Z"/>`),
		g.Group(children),
	)
}