package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opentasks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.67 6.09L42.5 11L21.59 31.89L11 21.31l4.82-4.82l5.75 5.74ZM28.48 25h10.34A1.24 1.24 0 0 1 40 26.14v4.6a1.24 1.24 0 0 1-1.15 1.15H6.65a1.15 1.15 0 0 1-1.15-1.15v-4.6A1.24 1.24 0 0 1 6.65 25h8m-8 10h32.17A1.23 1.23 0 0 1 40 36.17v4.59a1.23 1.23 0 0 1-1.15 1.15H6.65a1.15 1.15 0 0 1-1.15-1.15v-4.59A1.23 1.23 0 0 1 6.65 35Z"/>`),
		g.Group(children),
	)
}