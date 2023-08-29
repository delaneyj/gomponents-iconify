package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm2 3v12h14V10zm2.813 2h8.374L16 14.781zM11 13.875l4.438 2.969l.562.343l.563-.343L21 13.875V20H11z"/>`),
		g.Group(children),
	)
}