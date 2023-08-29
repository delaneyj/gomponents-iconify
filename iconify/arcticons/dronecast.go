package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dronecast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.74 30.7L6.5 42.43L24 5.57m4.26 25.13L41.5 42.43L24 5.57"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.74 30.7h8.52L24 5.57Z"/>`),
		g.Group(children),
	)
}