package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartThreeBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M6.444 3.685a10 10 0 0 1 1.662-.896c1.403-.593 2.104-.89 3-.296C12 3.086 12 4.057 12 6v2c0 1.886 0 2.828.586 3.414C13.172 12 14.114 12 16 12h2c1.942 0 2.914 0 3.507.895s.297 1.596-.296 3a10.002 10.002 0 0 1-11.162 5.913A10 10 0 0 1 6.444 3.685Z"/><path fill-rule="evenodd" d="M13.774 2.128a.75.75 0 0 1 .913-.54a10.768 10.768 0 0 1 7.725 7.725a.75.75 0 0 1-1.453.374a9.267 9.267 0 0 0-6.646-6.646a.75.75 0 0 1-.54-.913Z" clip-rule="evenodd" opacity=".5"/></g>`),
		g.Group(children),
	)
}