package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 18V2h-4v16h4zm-6 0V7H8v11h4zm-6 0v-8H2v8h4z"/>`),
		g.Group(children),
	)
}