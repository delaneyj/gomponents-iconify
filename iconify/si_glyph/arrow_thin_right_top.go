package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThinRightTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m16.908.021l.017 8.931c0 .648-.611 1.029-.895 1.027c-.57-.001-.948-.464-.949-1.034l-.01-5.557l-12.394 12.2a.94.94 0 0 1-1.328-.018c-.372-.372-.381-.966-.017-1.327L13.478 1.931l-5.333-.008c-.57-.002-1.062-.376-1.062-.944C7.08.409 7.568.003 8.14.005l8.768.016z"/>`),
		g.Group(children),
	)
}