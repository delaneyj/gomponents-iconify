package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinCornerArrowTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m511.9 232.7l-93.1-93V349L163 93.2h209.3L279.3.1L0 0l.1 279.3l93.1 93V163L349 418.8H139.7l93 93.1l279.3.1z"/>`),
		g.Group(children),
	)
}