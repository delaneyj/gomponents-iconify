package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 16.5l2-3l2 3m-2-3V12m2.627-1.376L15 10.5m-6 0l2.231.744"/><path d="M20.042 16.045A9 9 0 0 0 7.955 3.958M5.637 5.635a9 9 0 1 0 12.725 12.73"/><path d="M12 8a.5.5 0 1 0-.5-.5M3 3l18 18"/></g>`),
		g.Group(children),
	)
}