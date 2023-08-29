package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayListCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 6H3v2h12V6Zm0 4H3v2h12v-2ZM3 14h8v2H3v-2Zm8.99 1.025l1.415-1.414l2.121 2.121l4.243-4.242l1.414 1.414l-5.657 5.657l-3.535-3.536Z"/>`),
		g.Group(children),
	)
}