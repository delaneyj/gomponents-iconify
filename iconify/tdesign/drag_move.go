package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragMove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 23.414L7.586 19L9 17.586l2 2V13H4.414l2 2L5 16.414L.586 12L5 7.586L6.414 9l-2 2H11V4.414l-2 2L7.586 5L12 .586L16.414 5L15 6.414l-2-2V11h6.586l-2-2L19 7.586L23.414 12L19 16.414L17.586 15l2-2H13v6.586l2-2L16.414 19L12 23.414Z"/>`),
		g.Group(children),
	)
}