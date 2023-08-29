package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToriiGate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2.938 4.969V6.78l.312.313l4 3.656l.313.25H9v3H7v2h2v11h2V16h10v11h2V16h2v-2h-2v-3h1.375l.313-.25l4-3.656L29 6.78V4.97l-1.219.281C27.746 5.258 23.527 6 16 6C8.473 6 4.191 5.258 4.156 5.25zM6.78 7.562C8.895 7.774 11.914 8 16 8a91 91 0 0 0 9.156-.438L23.594 9H8.344zM11 11h3v3h-3zm7 0h3v3h-3z"/>`),
		g.Group(children),
	)
}