package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M5 9.5V19a3 3 0 0 1-.6 1.8L3.5 22M19 9.5V19a3 3 0 0 0 .6 1.8l.9 1.2M6.341 7A9.193 9.193 0 0 0 6 9.5c0 4.142 2.686 7.5 6 7.5s6-3.358 6-7.5C18 5.358 15.314 2 12 2c-1.09 0-2.114.364-2.995 1M5 20h7m7 0h-3"/><path d="M13 5.256c.96.51 1.697 1.732 1.926 3.244"/></g>`),
		g.Group(children),
	)
}