package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChartRing0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M43.776 20.994c-1.303-8.639-8.13-15.466-16.768-16.77m-6.032.003C11.366 5.685 4 13.982 4 24c0 10.02 7.37 18.32 16.986 19.774a20.165 20.165 0 0 0 6.018.002C35.646 42.474 42.476 35.643 43.776 27"/><path fill="#555" d="M24 16a8 8 0 1 0 0 16a8 8 0 0 0 0-16Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChartRing0)"/>`),
		g.Group(children),
	)
}