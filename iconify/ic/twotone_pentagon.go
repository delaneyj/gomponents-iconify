package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotonePentagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.63 9.78L16.56 19H7.44L4.37 9.78L12 4.44z" opacity=".3"/><path fill="currentColor" d="M19.63 9.78L16.56 19H7.44L4.37 9.78L12 4.44l7.63 5.34zM2 9l4 12h12l4-12l-10-7L2 9z"/>`),
		g.Group(children),
	)
}