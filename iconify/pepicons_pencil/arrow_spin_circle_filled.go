package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSpinCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilArrowSpinCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M16.374 8.038a.5.5 0 0 1-.563.827A5 5 0 1 0 18 13a.5.5 0 0 1 1 0a6 6 0 1 1-2.626-4.962Z"/><path d="M15.769 14.585a.5.5 0 1 1-.539-.842l3.482-2.227a.5.5 0 1 1 .539.842l-3.482 2.227Z"/><path d="M20.947 15.114a.5.5 0 0 1-.913.407l-1.509-3.38a.5.5 0 1 1 .914-.408l1.508 3.381Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilArrowSpinCircleFilled0)"/></g>`),
		g.Group(children),
	)
}