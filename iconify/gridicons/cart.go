package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20c0 1.1-.9 2-2 2s-1.99-.9-1.99-2S5.9 18 7 18s2 .9 2 2zm8-2c-1.1 0-1.99.9-1.99 2s.89 2 1.99 2s2-.9 2-2s-.9-2-2-2zm.396-5a2 2 0 0 0 1.952-1.566L21 5H7V4a2 2 0 0 0-2-2H3v2h2v11a2 2 0 0 0 2 2h12a2 2 0 0 0-2-2H7v-2h10.396z"/>`),
		g.Group(children),
	)
}