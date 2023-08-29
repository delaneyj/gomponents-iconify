package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DbnavigatorAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 31.5v-16h2.6a8.024 8.024 0 0 1 8 8h0a8.024 8.024 0 0 1-8 8Zm22.023-8.16a3.762 3.762 0 1 1 0 7.525h-6.208v-15.05h6.208a3.763 3.763 0 0 1 0 7.525Zm0 0h-6.208"/>`),
		g.Group(children),
	)
}