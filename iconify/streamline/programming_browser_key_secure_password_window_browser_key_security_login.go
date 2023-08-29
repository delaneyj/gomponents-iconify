package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingBrowserKeySecurePasswordWindowBrowserKeySecurityLogin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.44" cy="11.33" r="2.17"/><path d="m8 9.8l3.86-3.86a.36.36 0 0 1 .51 0l1.13 1.15m-3.05.28l1.02 1.02M2 12.5h-.5a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1V4m-13-.5h13"/></g>`),
		g.Group(children),
	)
}