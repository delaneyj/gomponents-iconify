package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oandbackupx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 30.07L24 37.52l19.5-7.45"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 25.26L24 32.71l19.5-7.45"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30 15.32L24 13L4.5 20.49L24 27.94l19.5-7.45m-10.08-9.96v10.08"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.84 17l-9.42 3.6L24 17"/>`),
		g.Group(children),
	)
}