package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataCollection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="16" r="2" fill="currentColor"/><path fill="currentColor" d="M30 17v-2h-6.17l2.58-2.59L25 11l-5 5l5 5l1.41-1.41L23.83 17H30zm-15 6.83V30h2v-6.17l2.59 2.58L21 25l-5-5l-5 5l1.41 1.41L15 23.83zM7 11l-1.41 1.41L8.17 15H2v2h6.17l-2.58 2.59L7 21l5-5l-5-5zm10-2.83V2h-2v6.17l-2.59-2.58L11 7l5 5l5-5l-1.41-1.41L17 8.17z"/>`),
		g.Group(children),
	)
}