package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 10v6H7v-6H2l8-8l8 8h-5zM0 18h20v2H0v-2z"/>`),
		g.Group(children),
	)
}