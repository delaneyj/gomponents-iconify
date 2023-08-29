package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20h11v6.17l-2.59-2.58L11 25l5 5l5-5l-1.41-1.41L17 26.17V20h11v-2H4v2zm7-13l1.41 1.41L15 5.83V12H4v2h24v-2H17V5.83l2.59 2.58L21 7l-5-5l-5 5z"/>`),
		g.Group(children),
	)
}