package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modalis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24.353" cy="24.05" r="9.578" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 6.81c16.13-.12 16.13-.211 18.147 5.142M42.5 41.19c-16.13.12-16.13.211-18.147-5.142"/>`),
		g.Group(children),
	)
}