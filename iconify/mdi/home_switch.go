package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 15v-4h3v4H13V9h2L8 3L1 9h2v6h3.5M9 16v2h6v-2l3 3l-3 3v-2H9v2l-3-3l3-3m14-7h-2v6h-6v-5h4l-5.46-4.89L16 3l7 6Z"/>`),
		g.Group(children),
	)
}