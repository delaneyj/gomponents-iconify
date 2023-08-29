package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 4h22v5.469l-.57.271a2.5 2.5 0 0 0 0 4.52l.57.271V20H1v-5.469l.57-.271a2.5 2.5 0 0 0 0-4.52L1 9.47V4Zm2 2v2.258c1.205.806 2 2.18 2 3.742a4.496 4.496 0 0 1-2 3.742V18h18v-2.258A4.496 4.496 0 0 1 19 12c0-1.561.795-2.936 2-3.742V6H3Zm5 3h8v2H8V9Zm0 4h8v2H8v-2Z"/>`),
		g.Group(children),
	)
}