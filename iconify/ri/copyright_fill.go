package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyrightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.52 0 10 4.48 10 10s-4.48 10-10 10S2 17.52 2 12S6.48 2 12 2Zm0 5c-2.76 0-5 2.24-5 5s2.24 5 5 5a5 5 0 0 0 4.288-2.428l-1.715-1.028A3 3 0 1 1 12 9c1.093 0 2.05.584 2.573 1.457l1.715-1.03A5 5 0 0 0 12 7Z"/>`),
		g.Group(children),
	)
}