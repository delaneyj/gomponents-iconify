package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChoppingBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChoppingBoard0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M14 8h28.011A1.99 1.99 0 0 1 44 10v28c0 1.105-.883 2-1.987 2H14c-4 0-10-2-10-16S11 8 14 8Z"/><path d="M12 20v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChoppingBoard0)"/>`),
		g.Group(children),
	)
}