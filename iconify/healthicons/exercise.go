package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exercise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M32 14h4v20h-4v-9H16v9h-4V14h4v9h16v-9ZM6 23v-6h4v14H6v-6H4v-2h2Zm38 2h-2v6h-4V17h4v6h2v2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}