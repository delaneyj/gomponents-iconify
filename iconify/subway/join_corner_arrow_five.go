package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinCornerArrowFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m512 512l-23.3-302.5L384 314.2L197.8 128L302.5 23.3L0 0l23.3 302.5L128 197.8L314.2 384L209.5 488.7z"/>`),
		g.Group(children),
	)
}