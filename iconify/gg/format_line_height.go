package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLineHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.097 6.997h2.077l-3-3l-3 3h1.923v10.006H1.159l3 3l3-3H5.097V6.997ZM22.841 7h-14V5h14v2Zm0 4h-14V9h14v2Zm-14 4h14v-2h-14v2Zm14 4h-14v-2h14v2Z"/>`),
		g.Group(children),
	)
}