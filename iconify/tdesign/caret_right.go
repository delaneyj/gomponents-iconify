package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 21.414L17.914 12L8.5 2.586v18.828Zm2-4.828V7.414L15.086 12L10.5 16.586Z"/>`),
		g.Group(children),
	)
}