package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApNews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.755 33.62V14.38h6.298a6.461 6.461 0 0 1 0 12.923h-6.298M9.5 33.562l6.385-19.181m6.12 19.238l-6.12-19.238m4.072 12.803h-8.334"/>`),
		g.Group(children),
	)
}