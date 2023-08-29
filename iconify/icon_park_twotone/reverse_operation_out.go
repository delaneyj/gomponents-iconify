package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReverseOperationOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTReverseOperationOut0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 5h18v32H6zm18 6h18v32H24z"/><path d="m17 17l-4 3.79L16.667 25M31 23l4 3.79L31.333 31"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTReverseOperationOut0)"/>`),
		g.Group(children),
	)
}