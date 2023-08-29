package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterXx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 9h-2l-2 6l-2-6H7l2.75 7L7 23h2l2-6l2 6h2l-2.76-7L15 9zm10 4h-2l-2 3.9l-2-3.9h-2l2.91 5L17 23h2l2-3.8l2 3.8h2l-2.9-5l2.9-5z"/>`),
		g.Group(children),
	)
}