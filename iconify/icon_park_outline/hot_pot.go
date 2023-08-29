package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotPot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 11V4m8 7V7m-16 4V7m28 10H4c0 5.586 3.578 10.503 9 13.365C16.156 32.03 19.936 33 24 33s7.844-.97 11-2.635c5.422-2.862 9-7.78 9-13.365ZM10.467 39h27.066M13 30.365L9 44m26-13.635L39 44M20 25h8"/>`),
		g.Group(children),
	)
}