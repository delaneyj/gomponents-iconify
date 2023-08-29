package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2h-2a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1H5c-1.103 0-2 .897-2 2v15zM5 5h2v2h10V5h2v15H5V5z"/><path fill="currentColor" d="M14.292 10.295L12 12.587l-2.292-2.292l-1.414 1.414l2.292 2.292l-2.292 2.292l1.414 1.414L12 15.415l2.292 2.292l1.414-1.414l-2.292-2.292l2.292-2.292z"/>`),
		g.Group(children),
	)
}