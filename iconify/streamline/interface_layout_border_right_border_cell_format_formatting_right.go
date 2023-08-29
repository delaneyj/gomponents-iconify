package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLayoutBorderRightBorderCellFormatFormattingRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 13.5h-1a1 1 0 0 1-1-1v-1m9 2h-4m-5-5v-3m13 8V.5M7 8.5v-3M5.5 7h3m-8 0h2m0-6.5h-1a1 1 0 0 0-1 1v1m9-2h-4m1.5 13v-2m0-9v-2"/>`),
		g.Group(children),
	)
}