package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stadtmobil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40 24L8 5.5A37.14 37.14 0 0 1 12.93 24A37.14 37.14 0 0 1 8 42.5Z"/>`),
		g.Group(children),
	)
}