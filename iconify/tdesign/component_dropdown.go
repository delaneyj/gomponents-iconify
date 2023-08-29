package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentDropdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v7H2V2Zm2 2v3h16V4H4Zm0 7v9h16v-9h2v11H2V11h2Zm2 1h12v2H6v-2Zm0 4h12v2H6v-2Z"/>`),
		g.Group(children),
	)
}