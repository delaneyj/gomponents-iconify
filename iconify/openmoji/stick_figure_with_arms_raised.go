package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickFigureWithArmsRaised(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m29.5 63.5l3-28.5h3M41 63.5L38 35h-2.5"/><path stroke-linecap="round" d="M32.5 35L30 19.5M38 35l2-15.5"/><circle cx="35" cy="11" r="3"/><path stroke-linecap="round" stroke-linejoin="round" d="m23 5l-1 10.5l8 4"/><path d="M30 19.5h10.5"/><path stroke-linecap="round" stroke-linejoin="round" d="m47 5l1 10.5l-8 4"/></g>`),
		g.Group(children),
	)
}