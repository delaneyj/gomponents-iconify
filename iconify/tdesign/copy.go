package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h13v5.5h-2V4H4v9h3.5v2H2V2Zm7 7h13v13H9V9Zm2 2v9h9v-9h-9Z"/>`),
		g.Group(children),
	)
}