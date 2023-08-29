package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1.5a5.5 5.5 0 0 1 3.352 9.86c1.888 1.05 3.148 2.96 3.148 5.14c0 3.314-2.91 6-6.5 6s-6.5-2.686-6.5-6c0-2.181 1.261-4.09 3.147-5.141A5.5 5.5 0 0 1 12 1.5Zm0 11c-2.52 0-4.5 1.828-4.5 4c0 2.172 1.98 4 4.5 4s4.5-1.828 4.5-4c0-2.172-1.98-4-4.5-4Zm0-9a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Z"/>`),
		g.Group(children),
	)
}