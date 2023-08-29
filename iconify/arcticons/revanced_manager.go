package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RevancedManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.386 4.5h21.109L23.94 27.241Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.316 4.5l-16.376 39l-16.256-39"/>`),
		g.Group(children),
	)
}