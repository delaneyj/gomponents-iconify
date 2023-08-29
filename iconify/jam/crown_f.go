package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrownF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.049 1.802L5.854 5.15L9.244.976a1 1 0 0 1 1.565.017l3.235 4.156l3.928-3.396a1 1 0 0 1 1.643.9L18.115 13H1.922L.399 2.7a1 1 0 0 1 1.65-.898zM2 14h16v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-1z"/>`),
		g.Group(children),
	)
}