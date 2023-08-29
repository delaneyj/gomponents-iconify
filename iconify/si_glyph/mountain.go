package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mountain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.012 1.062L4.035 8.87L2.709 7.569S-.305 14 .063 14h15.902v-.002l-3.338-7.1l-.983.612l-3.632-6.448zM5.611 7.521L8.062 2.77l2.285 4.081l-.986.67l-1.34-1.288l-2.41 1.288z"/>`),
		g.Group(children),
	)
}