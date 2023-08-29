package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storzbickel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.94 26.91H20.2l-6 16.31a21.64 21.64 0 0 1 0-38.44l6 16.31h7.62l6-16.31a21.64 21.64 0 0 1 0 38.44Z"/>`),
		g.Group(children),
	)
}