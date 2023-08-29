package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Form(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v3h16V4H4Zm16 5H4v11h16V9ZM6 11.5h12v2H6v-2Zm0 4h8v2H6v-2Z"/>`),
		g.Group(children),
	)
}