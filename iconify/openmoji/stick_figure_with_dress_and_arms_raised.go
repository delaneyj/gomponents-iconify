package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickFigureWithDressAndArmsRaised(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m29.5 63.5l3-28.5"/><path d="M29.5 63.5H41"/><path stroke-linecap="round" stroke-linejoin="round" d="M41 63.5L38 35"/><path d="M32.5 35L30 19m8 16l2-16"/><circle cx="35" cy="11" r="3"/><path stroke-linecap="round" stroke-linejoin="round" d="m23 5l-1 10.5l8 4M47 5l1 10.5l-8 4"/><path d="M30 19.5h10.5"/></g>`),
		g.Group(children),
	)
}