package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M150 56v120a6 6 0 0 1-12 0v-54H46v54a6 6 0 0 1-12 0V56a6 6 0 0 1 12 0v54h92V56a6 6 0 0 1 12 0Zm90 146h-36l36-47.95A30 30 0 1 0 187.71 126a6 6 0 1 0 11.29 4a18 18 0 0 1 14.47-11.82a18 18 0 0 1 16.87 28.66L187.2 204.4a6 6 0 0 0 4.8 9.6h48a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}