package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceiptSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.75 2A1.75 1.75 0 0 0 2 3.75v7.75A2.5 2.5 0 0 0 4.5 14h7a2.5 2.5 0 0 0 2.5-2.5V10h-3V3.75A1.75 1.75 0 0 0 9.25 2h-5.5ZM11 11h2v.5a1.5 1.5 0 0 1-1.5 1.5H11v-2ZM3 3.75A.75.75 0 0 1 3.75 3h5.5a.75.75 0 0 1 .75.75V13H4.5A1.5 1.5 0 0 1 3 11.5V3.75ZM4.5 5.5A.5.5 0 0 1 5 5h3a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5Zm0 2.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5Zm.5 2a.5.5 0 0 0 0 1h1.5a.5.5 0 0 0 0-1H5Z"/>`),
		g.Group(children),
	)
}