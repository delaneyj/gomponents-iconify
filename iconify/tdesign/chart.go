package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16V4H4Zm9 3v11h-2V7h2Zm-5 4v7H6v-7h2Zm10 2v5h-2v-5h2Z"/>`),
		g.Group(children),
	)
}