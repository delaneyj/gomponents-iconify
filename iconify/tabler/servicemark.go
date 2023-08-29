package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Servicemark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 9H6.5a1.5 1.5 0 0 0 0 3h1a1.5 1.5 0 0 1 0 3H5m8 0V9l3 4l3-4v6"/>`),
		g.Group(children),
	)
}