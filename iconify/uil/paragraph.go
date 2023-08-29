package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paragraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 13.5H3a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2Zm8-5H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}