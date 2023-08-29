package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLayoutBorderTopBorderCellFormatFormattingTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 11.5v1a1 1 0 0 1-1 1h-1m2-9v4m-5 5h-3m8-13H.5m8 6.5h-3M7 8.5v-3m0 8v-2m-6.5 0v1a1 1 0 0 0 1 1h1m-2-9v4m13-1.5h-2m-9 0h-2"/>`),
		g.Group(children),
	)
}