package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartThreeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M6.444 3.685a10 10 0 0 1 1.662-.896c1.403-.593 2.104-.89 3-.296C12 3.086 12 4.057 12 6v2c0 1.886 0 2.828.586 3.414C13.172 12 14.114 12 16 12h2c1.942 0 2.914 0 3.507.895s.297 1.596-.296 3a10.002 10.002 0 0 1-11.162 5.913A10 10 0 0 1 3.34 7M14.5 2.315A10.02 10.02 0 0 1 21.685 9.5"/>`),
		g.Group(children),
	)
}