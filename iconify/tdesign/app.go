package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func App(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h9v9H2V2Zm2 2v5h5V4H4Zm13.5 0a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM13 6.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0ZM2 13h9v9H2v-9Zm2 2v5h5v-5H4Zm9-2h9v9h-9v-9Zm2 2v5h5v-5h-5Z"/>`),
		g.Group(children),
	)
}