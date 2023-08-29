package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9.25 6.12V4.87h-7l3-2.05l-.71-1L.59 4.45A1.29 1.29 0 0 0 0 5.5a1.29 1.29 0 0 0 .59 1l3.93 2.72l.71-1l-3-2.06zm6.16 3.33l-3.93-2.67l-.71 1l3 2.06h-7v1.25h7.03l-3 2l.71 1l3.93-2.67a1.23 1.23 0 0 0 0-2.1z"/>`),
		g.Group(children),
	)
}