package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicRestEighth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 6a5.56 5.56 0 0 1-3.05 1.86A1.5 1.5 0 1 0 9.5 9h.24a6.32 6.32 0 0 0 3.51-1.07L10 18h2l4-12Z"/>`),
		g.Group(children),
	)
}