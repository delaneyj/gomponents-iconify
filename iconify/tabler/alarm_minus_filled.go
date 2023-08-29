package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmMinusFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M16 6.072a8 8 0 1 1-11.995 7.213L4 13l.005-.285A8 8 0 0 1 16 6.072zM14 12h-4l-.117.007A1 1 0 0 0 10 14h4l.117-.007A1 1 0 0 0 14 12z"/><path fill="currentColor" d="M6.412 3.191A1 1 0 0 1 7.685 4.73l-.097.08l-2.75 2a1 1 0 0 1-1.273-1.54l.097-.08l2.75-2zm9.779.221a1 1 0 0 1 1.291-.288l.106.067l2.75 2a1 1 0 0 1-1.07 1.685l-.106-.067l-2.75-2a1 1 0 0 1-.22-1.397z"/></g>`),
		g.Group(children),
	)
}