package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlignHorizontalLeftAlignDesignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5.5v13"/><rect width="6.5" height="4" x="3" y="1.5" rx="1"/><rect width="10.5" height="4" x="3" y="8.5" rx="1"/></g>`),
		g.Group(children),
	)
}