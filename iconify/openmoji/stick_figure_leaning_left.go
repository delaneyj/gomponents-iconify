package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickFigureLeaningLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M29.5 63.5L30 35l6 1.5m7.5 27l-7.5-27"/><path d="M30 35V19m6 17.5L39 18"/><path stroke-linecap="round" stroke-linejoin="round" d="m30 19l9.5-2l2 23M30 19l-5 8l5 7.5"/><circle cx="35" cy="11" r="3"/></g>`),
		g.Group(children),
	)
}