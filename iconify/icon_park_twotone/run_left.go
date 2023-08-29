package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRunLeft0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 5h18v32H6zm18 6h18v32H24z"/><path d="m17 17l-4 3.79L16.667 25M35 23l-4 3.79L34.667 31"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRunLeft0)"/>`),
		g.Group(children),
	)
}