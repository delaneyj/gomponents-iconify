package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.25 37.45v-6.17s-8-7.44-8-15.86S18.11 4.5 24 4.5S36.81 7 36.81 15.42s-8.06 15.86-8.06 15.86v6.17m-9.5 0h9.5l-9.5 6.05h9.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.03 18.61l-6.71-6.71l1.36 12.06l-6.71-6.72"/>`),
		g.Group(children),
	)
}