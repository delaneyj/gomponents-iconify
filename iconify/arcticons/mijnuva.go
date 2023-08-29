package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mijnuva(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.5 19.5l9 9m0-9l-9 9m0-23l9 9m0-9l-9 9m0 19l9 9m0-9l-9 9m-4.218-24.864v7.343c0 5.386 1.56 7 7 7c4.176 0 10.436 0 10.436-4.25m0 4.25V17.636"/>`),
		g.Group(children),
	)
}