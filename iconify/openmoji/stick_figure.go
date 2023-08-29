package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickFigure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m29.5 63.5l3-28.5h3M41 63.5L38 35h-2.5"/><path d="M32.5 35L30 18m8 17l2.5-17"/><path stroke-linecap="round" stroke-linejoin="round" d="M35 18h5.5L43 42m-8-24h-5l-2 24"/><circle cx="35" cy="11" r="3"/></g>`),
		g.Group(children),
	)
}