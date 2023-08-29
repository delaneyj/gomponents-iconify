package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingBrowserDashboardAppCodeAppsProgrammingWindowDashboardPerformance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M.5 3.5h13m-6.5 7L10.5 7m0 3.5h1M7 7V6M4.53 8.03l-.71-.71M3.5 10.5h-1"/></g>`),
		g.Group(children),
	)
}