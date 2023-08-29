package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.83 8.79L8 9.456V13H6V8.05h.015l5.268-1.918c.244-.093.51-.14.782-.131a2.616 2.616 0 0 1 2.427 1.82c.186.583.356.977.51 1.181A4.992 4.992 0 0 0 19 11v2a6.986 6.986 0 0 1-5.402-2.547l-.697 3.956L15 16.17V23h-2v-5.898l-2.27-1.904l-.727 4.127l-6.894-1.215l.348-1.97l4.924.868L9.83 8.79ZM13.5 5.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}