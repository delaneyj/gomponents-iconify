package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingBrowserSearchSearchWindowGlassAppCodeProgrammingQueryFindMagnifyingApps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 12.5h-4a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v5m-13-3h13"/><circle cx="9.5" cy="9.5" r="2.5"/><path d="m11.27 11.27l2.23 2.23"/></g>`),
		g.Group(children),
	)
}