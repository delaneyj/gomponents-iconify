package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbTestingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1.5 11V6.5a2 2 0 1 1 4 0V11m-4-2.5h4m6.5-1H9.5m2.5 0a1.5 1.5 0 0 0 0-3H9.5v3m2.5 0a1.5 1.5 0 0 1 0 3H9.5v-3M7.5 1v13"/>`),
		g.Group(children),
	)
}