package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemixiconLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.365 6l8.784 9.663l.72-.283c1.685-.661 2.864-2.156 3.092-3.896A6.502 6.502 0 0 1 12.078 6H6.364ZM14 5a4.5 4.5 0 0 0 6.714 3.918c.186.618.286 1.271.286 1.947c0 2.891-1.822 5.365-4.4 6.377L20 21H3V4h11.111A4.512 4.512 0 0 0 14 5Zm4.5 2.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5ZM5 7.47V19h10.48L5.001 7.47Z"/>`),
		g.Group(children),
	)
}