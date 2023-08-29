package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cflumen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.5 24A21.494 21.494 0 0 1 32 4.055a21.5 21.5 0 1 0 0 39.89A21.494 21.494 0 0 1 18.5 24Zm20.326-6.674l2.36 4.314L45.5 24l-4.314 2.36l-2.36 4.314l-2.359-4.314L32.152 24l4.315-2.36l2.359-4.314z"/>`),
		g.Group(children),
	)
}