package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscountFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.474 7.524a1.5 1.5 0 1 0-2.122 2.122a1.5 1.5 0 0 0 2.122-2.122Z"/><path fill="currentColor" d="m22.003 1.998l.808 10.505l-10.932 10.932L.565 12.121L11.497 1.19l10.506.809Zm-8.711 4.466a3 3 0 1 0 4.243 4.242a3 3 0 0 0-4.243-4.242Z"/>`),
		g.Group(children),
	)
}