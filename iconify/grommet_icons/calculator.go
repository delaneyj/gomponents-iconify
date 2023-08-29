package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M22 23H2V1h20v22Zm-6-6h2v2h-2v-2Zm-5 0h2v2h-2v-2Zm5-5h2v2h-2v-2Zm-5 0h2v2h-2v-2Zm-5 5h2v2H6v-2Zm0-5h2v2H6v-2Zm12-3H6V5h12v4Z"/>`),
		g.Group(children),
	)
}