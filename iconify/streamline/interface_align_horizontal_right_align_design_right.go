package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlignHorizontalRightAlignDesignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5.5v13"/><rect width="6.5" height="4" x="4.5" y="1.5" rx="1" transform="rotate(180 7.75 3.5)"/><rect width="10.5" height="4" x=".5" y="8.5" rx="1" transform="rotate(180 5.75 10.5)"/></g>`),
		g.Group(children),
	)
}