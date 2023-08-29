package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickFigureWithDressLeaningRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m42 35l.5 28.5h-14m0 0l7.5-27"/><path d="M42 35V19m-6 17.5L33 18"/><path stroke-linecap="round" stroke-linejoin="round" d="m42 19l-9.5-2l-2 23M42 19l5 8l-5 7.5"/><circle r="3" transform="matrix(-1 0 0 1 37 11)"/></g>`),
		g.Group(children),
	)
}