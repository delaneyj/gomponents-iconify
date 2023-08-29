package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLayoutBorderFrameBorderCellFormatFormattingFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 5.5v3M8.5 7h-3m8 0h-2m-9 0h-2"/><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M7 .5v2m0 9v2"/></g>`),
		g.Group(children),
	)
}