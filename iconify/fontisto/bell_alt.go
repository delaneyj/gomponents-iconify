package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.333 17.16c-1.04-1.04-2.339-2.341-2.339-7.409a7.505 7.505 0 0 0-6.22-7.393l-.045-.006a1.5 1.5 0 1 0-2.467.005l-.003-.005c-3.578.614-6.266 3.692-6.266 7.399c0 5.068-1.296 6.367-2.339 7.409A2.252 2.252 0 0 0 2.245 21h5.249a3 3 0 1 0 6 0h5.248a2.253 2.253 0 0 0 1.591-3.84zm-9.84 4.965a.375.375 0 0 1 0 .75A1.877 1.877 0 0 1 8.618 21h.75a1.127 1.127 0 0 0 1.126 1.125z"/>`),
		g.Group(children),
	)
}