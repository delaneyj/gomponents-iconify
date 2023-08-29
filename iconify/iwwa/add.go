package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Add(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="m36.495 19.226l-15.732.012l.012-15.732a.763.763 0 0 0-1.525 0l-.012 15.733l-15.733.011a.763.763 0 0 0 0 1.525l15.732-.012l-.012 15.732c0 .204.082.4.223.538a.764.764 0 0 0 1.303-.538l.012-15.732l15.732-.012a.763.763 0 1 0 0-1.525z"/>`),
		g.Group(children),
	)
}