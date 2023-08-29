package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextItalic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 9V7H12v2h5.14l-4.37 14H7v2h13v-2h-5.14l4.37-14H25z"/>`),
		g.Group(children),
	)
}