package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bibleinoneyear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.491 4.5v39L24 36.241L38.509 43.5v-39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.81 8.13h-9.069v22.672h9.07c9.669 0 10.168-12.698.913-12.698c5.93 0 5.93-9.552-.914-9.975Z"/>`),
		g.Group(children),
	)
}