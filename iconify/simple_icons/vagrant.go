package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vagrant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.556 0L.392 1.846V4.11l7.124 17.3L11.998 24l4.523-2.611l7.083-17.345V1.848l.004-.002L20.44 0l-5.274 3.087v2.111l-3.168 7.384l-3.164-7.384V3.109l-.017-.008l.017-.01z"/>`),
		g.Group(children),
	)
}