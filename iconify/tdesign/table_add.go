package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v2h16V4H4Zm16 4H4v12h16V8Zm-7 2v3h3v2h-3v3h-2v-3H8v-2h3v-3h2Z"/>`),
		g.Group(children),
	)
}