package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vrv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.818 44.633L5.84 34.308l.088-20.5L23.994 3.633L41.973 13.96l-.088 20.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.883 28.072l14.273 8.24L39.1 27.685"/>`),
		g.Group(children),
	)
}