package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19.184 4.666L12 4.5l-7.184.166A3.057 3.057 0 0 0 1.87 7.17a26.718 26.718 0 0 0 0 9.66a3.058 3.058 0 0 0 2.945 2.504L12 19.5l7.184-.166a3.058 3.058 0 0 0 2.945-2.504a26.717 26.717 0 0 0 0-9.66a3.057 3.057 0 0 0-2.945-2.504ZM21 11a1 1 0 0 1-1 1H4a1 1 0 1 1 0-2h16a1 1 0 0 1 1 1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}