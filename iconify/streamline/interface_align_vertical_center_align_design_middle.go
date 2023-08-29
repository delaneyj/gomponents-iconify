package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlignVerticalCenterAlignDesignMiddle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="4" x="7.5" y="5.5" rx="1" transform="rotate(90 11.5 7.5)"/><rect width="11" height="4" x="-3" y="5" rx="1" transform="rotate(90 2.5 7)"/><path d="M9.5 7h-5"/></g>`),
		g.Group(children),
	)
}