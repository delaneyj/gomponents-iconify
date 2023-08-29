package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExchangeFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSExchangeFour0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke="#000" d="M15 21h18m-18 6h18m0-6l-7-7m-4 20l-7-7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSExchangeFour0)"/>`),
		g.Group(children),
	)
}