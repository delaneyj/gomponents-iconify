package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Macys(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.015.624L9.19 9.293H0l7.445 5.384l-2.819 8.673L12 17.986l7.422 5.393l-2.835-8.713L24 9.292h-9.162L12.015.622v.002z"/>`),
		g.Group(children),
	)
}