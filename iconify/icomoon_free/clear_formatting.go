package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearFormatting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 14h9v2H0zM14 2H9.273L6.402 13H4.335L7.206 2H3.001V0h11zm.528 14L12.5 13.972L10.472 16l-.972-.972L11.528 13L9.5 10.972l.972-.972l2.028 2.028L14.528 10l.972.972L13.472 13l2.028 2.028z"/>`),
		g.Group(children),
	)
}