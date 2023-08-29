package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TumblrSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2v20h20V2Zm11.57 15.85a3.4 3.4 0 0 1-3.75-3.33v-4.08H8.56V8.83A3.83 3.83 0 0 0 11 5.47a.09.09 0 0 1 .09-.09h1.82v3.17h2.5v1.89h-2.49v3.89a1.11 1.11 0 0 0 1.2 1.23a3 3 0 0 0 1-.22l.6 1.78a3.34 3.34 0 0 1-2.15.73Z"/>`),
		g.Group(children),
	)
}