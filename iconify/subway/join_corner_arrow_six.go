package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinCornerArrowSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M209.5 23.3L314.2 128L128 314.2L23.3 209.5L0 512l302.5-23.3L197.8 384L384 197.8l104.7 104.7L512 0z"/>`),
		g.Group(children),
	)
}