package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmSnoozeFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M16 6.072a8 8 0 1 1-11.995 7.213L4 13l.005-.285A8 8 0 0 1 16 6.072zM14 10h-4l-.117.007A1 1 0 0 0 9 11l.007.117A1 1 0 0 0 10 12h1.584l-2.291 2.293l-.076.084C8.703 15.014 9.147 16 10 16h4l.117-.007A1 1 0 0 0 15 15l-.007-.117A1 1 0 0 0 14 14h-1.586l2.293-2.293l.076-.084c.514-.637.07-1.623-.783-1.623z"/><path fill="currentColor" d="M6.412 3.191A1 1 0 0 1 7.685 4.73l-.097.08l-2.75 2a1 1 0 0 1-1.273-1.54l.097-.08l2.75-2zm9.779.221a1 1 0 0 1 1.291-.288l.106.067l2.75 2a1 1 0 0 1-1.07 1.685l-.106-.067l-2.75-2a1 1 0 0 1-.22-1.397z"/></g>`),
		g.Group(children),
	)
}