package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleObtuse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 19H9.31L4.07 5.36l1.86-.72l3.03 7.86a4.82 4.82 0 0 1 1.75-.33A4.88 4.88 0 0 1 15.58 17H21v2m-10.31-2h2.89a2.849 2.849 0 0 0-2.87-2.83c-.37 0-.71.07-1.04.19L10.69 17Z"/>`),
		g.Group(children),
	)
}