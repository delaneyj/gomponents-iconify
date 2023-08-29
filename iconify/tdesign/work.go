package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Work(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 2.5h9v4H22v15H2v-15h5.5v-4Zm2 4h5v-2h-5v2ZM4 8.5v11h16v-11H4Z"/>`),
		g.Group(children),
	)
}