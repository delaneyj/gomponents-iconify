package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m11.75 5l-.063.938l-.187 3L11.437 10h2.001l-.876 12h-2l-.062.938l-.188 3L10.22 27h10.03l.063-.938l.187-3L20.563 22h-2.001l.875-12h2.001l.062-.938l.188-3L21.78 5H11.75zm1.875 2h6l-.063 1h-2l-.062.938l-1 14L16.437 24h2.001l-.063 1h-6l.063-1h2l.062-.938l1-14L15.563 8h-2.001l.063-1z"/>`),
		g.Group(children),
	)
}