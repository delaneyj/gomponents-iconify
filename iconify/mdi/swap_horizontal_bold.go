package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 10v3h6v5H8v3l-6-5.5L8 10m14-1.5L16 3v3h-6v5h6v3l6-5.5Z"/>`),
		g.Group(children),
	)
}