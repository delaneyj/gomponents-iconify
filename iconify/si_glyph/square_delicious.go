package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareDelicious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 0v16h16V0H1zm7.938 15.016V8.01H1.959V.958H9.01v7.011h7.006v7.047H8.938z"/>`),
		g.Group(children),
	)
}