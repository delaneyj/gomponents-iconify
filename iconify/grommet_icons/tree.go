package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4 1h6v6H4V1Zm12 10h4v4h-4v-4Zm0 8h4v4h-4v-4ZM7 7v14h9m-9-8h9"/>`),
		g.Group(children),
	)
}