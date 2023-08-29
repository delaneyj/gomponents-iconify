package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7.587 7.566a7 7 0 1 0 9.833 9.864m1.35-2.645a7 7 0 0 0-8.536-8.56"/><path d="M12 12v1h1M5.261 5.265L4.25 6M17 4l2.75 2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}