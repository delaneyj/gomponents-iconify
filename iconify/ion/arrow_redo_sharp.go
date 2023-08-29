package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRedoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M48 399.26C48 335.19 62.44 284 90.91 247c34.38-44.67 88.68-68.77 161.56-71.75V72L464 252L252.47 432V329.35c-44.25 1.19-77.66 7.58-104.27 19.84c-28.75 13.25-49.6 33.05-72.08 58.7L48 440Z"/>`),
		g.Group(children),
	)
}