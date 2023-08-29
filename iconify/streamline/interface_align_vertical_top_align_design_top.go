package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlignVerticalTopAlignDesignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5.5H.5"/><rect width="6.5" height="4" x="7.25" y="4.25" rx="1" transform="rotate(90 10.5 6.25)"/><rect width="10.5" height="4" x="-1.75" y="6.25" rx="1" transform="rotate(90 3.5 8.25)"/></g>`),
		g.Group(children),
	)
}