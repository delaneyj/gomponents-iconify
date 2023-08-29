package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLayoutBorderLeftBorderCellFormatFormattingLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5.5h1a1 1 0 0 1 1 1v1m-9-2h4m5 5v3M.5.5v13m6.5-8v3M8.5 7h-3m8 0h-2m0 6.5h1a1 1 0 0 0 1-1v-1m-9 2h4M7 .5v2m0 9v2"/>`),
		g.Group(children),
	)
}