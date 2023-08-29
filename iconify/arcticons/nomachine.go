package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nomachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 39.376V8.624H13.562v30.752M27.96 8.624v30.752M5.5 8.624v23.611m0 4.146v2.995"/>`),
		g.Group(children),
	)
}