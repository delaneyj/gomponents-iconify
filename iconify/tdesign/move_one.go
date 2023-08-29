package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.086 12L5.5 7.586L6.914 9l-2 2H11V4.914l-2 2L7.586 5.5L12 1.086L16.414 5.5L15 6.914l-2-2V11h6.086l-2-2L18.5 7.586L22.914 12L18.5 16.414L17.086 15l2-2H13v6.086l2-2l1.414 1.414L12 22.914L7.586 18.5L9 17.086l2 2V13H4.914l2 2L5.5 16.414L1.086 12Z"/>`),
		g.Group(children),
	)
}