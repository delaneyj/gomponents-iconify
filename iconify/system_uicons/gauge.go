package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gauge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" transform="translate(2 3)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14 14c1.448-1.448 2.5-3.29 2.5-5.5a8 8 0 1 0-16 0c0 2.21 1.052 4.052 2.5 5.5m5.5-5.5l-4-4"/><circle cx="8.5" cy="8.5" r="1.5" fill="currentColor"/></g>`),
		g.Group(children),
	)
}