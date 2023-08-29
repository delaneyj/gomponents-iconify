package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterVv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 13h-1.75L21 22.03L18.79 13H17l2.5 10h3L25 13zM13 9l-2 13L9 9H7l2.52 14h2.96L15 9h-2z"/>`),
		g.Group(children),
	)
}