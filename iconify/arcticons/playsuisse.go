package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playsuisse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.78 21.06h-7.17v-7.17h-5.88v7.17h-7.17v5.88h7.17v7.17h5.88v-7.17h7.17v-5.88zM8.94 14.77l-4.16 4.16L9.85 24l-5.07 5.07l4.16 4.16L18.16 24l-9.22-9.23z"/>`),
		g.Group(children),
	)
}