package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Edpzero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.34 36.27A21.51 21.51 0 0 1 36.27 6.34m5.39 5.39a21.51 21.51 0 0 1-29.93 29.93"/>`),
		g.Group(children),
	)
}