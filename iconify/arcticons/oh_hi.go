package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OhHi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M7.5 5.5a1.996 1.996 0 0 0-2 2v13a1.996 1.996 0 0 0 2 2h13a1.996 1.996 0 0 0 2-2v-13a1.996 1.996 0 0 0-2-2Zm20 0a1.996 1.996 0 0 0-2 2v13a1.996 1.996 0 0 0 2 2h13a1.996 1.996 0 0 0 2-2v-13a1.996 1.996 0 0 0-2-2Zm-20 20a1.996 1.996 0 0 0-2 2v13a1.996 1.996 0 0 0 2 2h13a1.996 1.996 0 0 0 2-2v-13a1.996 1.996 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}