package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v7h7V4H4Zm9 0v7h7V4h-7Zm7 9h-7v7h7v-7Zm-9 7v-7H4v7h7Z"/>`),
		g.Group(children),
	)
}