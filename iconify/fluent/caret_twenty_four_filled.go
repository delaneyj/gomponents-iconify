package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M18 7.207c0-1.114-1.346-1.671-2.134-.884l-9.543 9.543c-.787.788-.23 2.134.884 2.134h9.043A1.75 1.75 0 0 0 18 16.25V7.207z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}