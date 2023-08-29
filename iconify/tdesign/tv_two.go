package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 3h24v15H0V3Zm2 2v11h20V5H2Zm2 15h16v2H4v-2Z"/>`),
		g.Group(children),
	)
}