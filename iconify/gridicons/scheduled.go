package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scheduled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.498 18.001l-3.705-3.705l1.415-1.415l2.294 2.294l5.293-5.293l1.415 1.415l-6.712 6.704zM21 6v13a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h1V2h2v2h8V2h2v2h1a2 2 0 0 1 2 2zm-2 2H5v11h14V8z"/>`),
		g.Group(children),
	)
}