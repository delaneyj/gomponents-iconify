package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighlightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 1a1 1 0 0 0-1 1v2.5A1.5 1.5 0 0 0 3.5 6c-.017 0-.033 0-.05.003h.102A.527.527 0 0 0 3.5 6h.246v.003h8.545V6h.21a.51.51 0 0 0-.051.003h.102A.506.506 0 0 0 12.5 6h.001a1.5 1.5 0 0 0 1.5-1.5V2a1 1 0 0 0-1-1H3Zm.001 6.003a2 2 0 0 0 2 1.997h6a2 2 0 0 0 2-1.997h-10ZM5.003 14.5V10H11l.002.74a1.5 1.5 0 0 1-.69 1.265l-4.54 2.916a.5.5 0 0 1-.77-.421Z"/>`),
		g.Group(children),
	)
}