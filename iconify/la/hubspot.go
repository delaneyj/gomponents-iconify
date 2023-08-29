package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hubspot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7.5 4a2.5 2.5 0 0 0 0 5a2.5 2.5 0 0 0 1.332-.389l7.486 5.641A5.97 5.97 0 0 0 15 18c0 1.552.594 2.96 1.56 4.025l-3.044 3.045A2 2 0 0 0 13 25a2 2 0 0 0 0 4a2 2 0 0 0 1.93-2.516l3.209-3.209h.002A6 6 0 0 0 27 18a5.997 5.997 0 0 0-5.001-5.91V8.73A2 2 0 0 0 21 5a2 2 0 0 0-1 3.73v3.36a5.955 5.955 0 0 0-2.15.812L9.957 6.955A2.5 2.5 0 0 0 7.5 4zM21 15c1.654 0 3 1.346 3 3s-1.346 3-3 3s-3-1.346-3-3s1.346-3 3-3z"/>`),
		g.Group(children),
	)
}