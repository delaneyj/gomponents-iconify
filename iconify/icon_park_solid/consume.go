package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Consume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSConsume0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 14a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V14Z"/><path stroke="#000" stroke-linecap="round" d="m19 19l5 5l5-5m-11 6h12m-12 6h12m-6-6v10"/><path stroke="#fff" stroke-linecap="round" d="M8 6h32"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSConsume0)"/>`),
		g.Group(children),
	)
}