package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mypat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.62 25.186a15.341 15.341 0 1 1 8.213 8.703"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.42 43.5s-.8-18.337 10.85-28.263"/>`),
		g.Group(children),
	)
}