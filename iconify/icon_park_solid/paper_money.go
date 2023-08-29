package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperMoney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPaperMoney0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M39 7H9a5 5 0 0 0-5 5v24a5 5 0 0 0 5 5h30a5 5 0 0 0 5-5V12a5 5 0 0 0-5-5Z"/><path stroke="#000" stroke-linecap="round" d="m18 15l6 6l6-6m-13 8h14m-14 6h14m-7-6v11"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPaperMoney0)"/>`),
		g.Group(children),
	)
}