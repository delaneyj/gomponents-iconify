package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Manman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.5v35a2 2 0 0 0 2 2h2.33v-39H10.4a2 2 0 0 0-2 2Zm4.331-2v39h24.87a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.448 32.961l-.002-2.751l-2.281-1.374l-2.28 1.376l.001 2.751m4.562-17.924l-.002 2.751l-2.281 1.374l-2.28-1.376l.001-2.751m2.279 13.798l-.003-9.67"/>`),
		g.Group(children),
	)
}