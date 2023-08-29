package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsExpandFourExpandSmallBiggerRetractSmallerBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="10" x="3.5" y=".5" rx="1" transform="rotate(180 8.5 5.5)"/><path d="M10.5 13.5h-9a1 1 0 0 1-1-1v-9m7 0h3v3m0-3l-4 4"/></g>`),
		g.Group(children),
	)
}