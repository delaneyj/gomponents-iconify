package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Consume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTConsume0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 14a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V14Z"/><path stroke-linecap="round" d="m19 19l5 5l5-5m-11 6h12m-12 6h12m-6-6v10M8 6h32"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTConsume0)"/>`),
		g.Group(children),
	)
}