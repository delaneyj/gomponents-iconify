package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7 3h15M7 9h9m-9 6h15M2 2h2v2H2V2Zm0 6h2v2H2V8Zm0 6h2v2H2v-2Zm0 6h2v2H2v-2Zm5 1h9"/>`),
		g.Group(children),
	)
}