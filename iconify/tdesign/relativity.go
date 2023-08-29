package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Relativity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h13v7h7v13H9v-7H2V2Zm9 13v5h9v-9h-5v4h-4Zm2-6V4H4v9h5V9h4Zm0 2h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}