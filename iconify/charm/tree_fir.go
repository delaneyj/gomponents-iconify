package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeFir(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 1.75l-4.25 5.5h2.5l-3.5 4h4v3h2.5v-3h4l-3.5-4h2.5z"/>`),
		g.Group(children),
	)
}