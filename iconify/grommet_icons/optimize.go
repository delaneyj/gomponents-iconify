package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Optimize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M2 22h4v-4H2v4ZM22 2L12 12m10-2V2h-8m8 11h-4v9h4v-9Zm-12 9h4v-6h-4v6Z"/>`),
		g.Group(children),
	)
}