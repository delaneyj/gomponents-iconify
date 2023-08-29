package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLayoutOneColumnLayoutLayoutsLeftSidebar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(180 7 7)"/><path d="M5.5.5v13m0-6.5h8"/></g>`),
		g.Group(children),
	)
}