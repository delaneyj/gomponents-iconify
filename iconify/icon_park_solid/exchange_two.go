package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExchangeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSExchangeTwo0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M9 18v24h30V18L24 6L9 18Z"/><path stroke="#000" d="M24 30v6m7-10v10m-14-4v4m0-11l5-4l3 3l6-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSExchangeTwo0)"/>`),
		g.Group(children),
	)
}