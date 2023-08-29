package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bird(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 68a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm-50.88 61.85a8 8 0 0 0-11.27 1l-40 48a8 8 0 0 0 12.3 10.24l40-48a8 8 0 0 0-1.03-11.24ZM240 80a8 8 0 0 1-3.56 6.66L216 100.28V120a104.11 104.11 0 0 1-104 104H8a8 8 0 0 1-6.25-13L96 93.19V76a60 60 0 0 1 117.21-18.14l23.23 15.48A8 8 0 0 1 240 80Zm-22.42 0L201.9 69.54a8 8 0 0 1-3.31-4.64A44 44 0 0 0 112 76v20a8 8 0 0 1-1.75 5L24.64 208H112a88.1 88.1 0 0 0 88-88V96a8 8 0 0 1 3.56-6.66Z"/>`),
		g.Group(children),
	)
}