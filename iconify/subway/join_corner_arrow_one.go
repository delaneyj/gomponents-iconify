package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinCornerArrowOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m232.7.1l-93 93.1H349L93.2 349V139.7l-93.1 93L0 512l279.3-.1l93-93.1H163L418.8 163v209.3l93.1-93L512 0z"/>`),
		g.Group(children),
	)
}