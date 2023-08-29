package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func G(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M33.49 16.39a2 2 0 0 1-2.823.156A10 10 0 1 0 30 32v-6h-6a2 2 0 1 1 0-4h8a2 2 0 0 1 2 2v8.944a2 2 0 0 1-.667 1.491a14 14 0 1 1 0-20.87a2 2 0 0 1 .158 2.824Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}