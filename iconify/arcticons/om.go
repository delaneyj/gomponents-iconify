package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Om(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.68 26.53H7a2.53 2.53 0 0 1 0-5.06h12.68a2.53 2.53 0 0 1 0 5.06Zm21.29 0H28.32a2.53 2.53 0 0 1 0-5.06H41a2.53 2.53 0 0 1 0 5.06Z"/>`),
		g.Group(children),
	)
}