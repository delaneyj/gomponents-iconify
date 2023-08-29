package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="M16.91 29.83a.997.997 0 0 1-.707-.293l-11.48-11.48l1.414-1.414L16.91 27.416l16.953-16.953l1.414 1.414l-17.66 17.66a.997.997 0 0 1-.707.293z"/>`),
		g.Group(children),
	)
}