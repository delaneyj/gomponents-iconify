package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Githo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24a21.478 21.478 0 1 0-1.01 6.529"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.453 32.444l5.038-1.915l1.915 5.037"/>`),
		g.Group(children),
	)
}