package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BalanceOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBalanceOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path fill="#555" d="M35 25a32.234 32.234 0 0 0-22 0l-1-11c7-3 17-3 24 0l-1 11Z"/><path d="m24 23l-3-5m21 21a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v30Z"/><path d="M29 23.455a32.228 32.228 0 0 0-10 0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBalanceOne0)"/>`),
		g.Group(children),
	)
}