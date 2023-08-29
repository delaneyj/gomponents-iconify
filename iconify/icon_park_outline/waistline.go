package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Waistline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 31c3.056 1.51 13.2 3.622 22 0"/><circle cx="24" cy="25.412" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 5c1.856 3.278 4.971 10.258 5 17c.009 2.065-.319 4.107-1 6.014C12.18 33.1 9.453 36.042 9 43M38 5c-1.856 3.278-4.971 10.258-5 17c-.009 2.065.319 4.107 1 6.014C35.82 33.1 38.546 36.042 39 43M8 18s1 2 1 4c0 1.61-1 4-1 4m32-8s-1 2-1 4c0 1.61 1 4 1 4"/></g>`),
		g.Group(children),
	)
}