package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flowx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m24 4.5l-12.604 18c-4.901 7-.396 21 12.604 21s17.505-14 12.604-21L24 4.5Z"/><path d="m24 13.5l-7 9.997C13.096 29.072 18 36.5 24 36.5s11.002-7.287 7.002-13L24 13.5Zm0 0l3-4.284"/><path d="M29.69 42.47c-6.041.486-13.372-5.64-14.054-13.974"/></g>`),
		g.Group(children),
	)
}