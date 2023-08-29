package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gasflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 9.61a9.15 9.15 0 0 1 9.14 9.14h0A9.15 9.15 0 0 1 24 27.88h0a9.14 9.14 0 1 1 0-18.27Zm4.57 4.57l-5.82 3.32l-.2 2.69l2.7-.19l3.32-5.82Z"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.07 32.04h39.86"/>`),
		g.Group(children),
	)
}