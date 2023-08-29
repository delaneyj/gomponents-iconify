package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feItalic0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feItalic1" fill="currentColor"><path id="feItalic2" d="m8.97 19l3.75-14H10a1 1 0 1 1 0-2h8a1 1 0 0 1 0 2h-3.208L11.04 19H14a1 1 0 0 1 0 2H6a1 1 0 0 1 0-2h2.97Z"/></g></g>`),
		g.Group(children),
	)
}