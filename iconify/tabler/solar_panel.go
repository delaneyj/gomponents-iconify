package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolarPanel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.28 14h15.44a1 1 0 0 0 .97-1.243l-1.5-6A1 1 0 0 0 18.22 6H5.78a1 1 0 0 0-.97.757l-1.5 6A1 1 0 0 0 4.28 14zM4 10h16M10 6l-1 8m5-8l1 8m-3 0v4m-5 0h10"/>`),
		g.Group(children),
	)
}