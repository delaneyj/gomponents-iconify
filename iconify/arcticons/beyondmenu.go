package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beyondmenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.12 36.24l7-24.38l9.31 13.82l-4.08 7.73l-4-7.68l9.26-14l7 24.43M4.64 24.24L4.5 35.91c17.32 0 13.12-17.2 5.09-17.2M4.77 12.7L15 12.57L4.64 24.24Z"/>`),
		g.Group(children),
	)
}