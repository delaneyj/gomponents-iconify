package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRotaryLastRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0m3 3v6m2.5-11.5L18 3m-5 0h5v5"/>`),
		g.Group(children),
	)
}