package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 0a1.2 1.2 0 0 0-1.2 1.2v10.399a5.2 5.2 0 1 1-10.4 0V1.2A1.202 1.202 0 0 0 0 1.198v10.401a7.6 7.6 0 0 0 15.2 0V1.2A1.2 1.2 0 0 0 14.001 0zm0 21.6H1.2a1.202 1.202 0 0 0-.002 2.4H14a1.202 1.202 0 0 0 .002-2.4z"/>`),
		g.Group(children),
	)
}