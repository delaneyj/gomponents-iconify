package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StethoscopeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.172 4.179A2 2 0 0 0 3 6v3.5a5.5 5.5 0 0 0 9.856 3.358M14 10V6a2 2 0 0 0-2-2h-1"/><path d="M8 15a6 6 0 0 0 10.714 3.712m1.216-2.798c.046-.3.07-.605.07-.914v-3m-9-9v2"/><path d="M18 10a2 2 0 1 0 4 0a2 2 0 1 0-4 0M3 3l18 18"/></g>`),
		g.Group(children),
	)
}