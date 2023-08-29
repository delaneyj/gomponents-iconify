package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullSelection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFullSelection0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M34 5H8a3 3 0 0 0-3 3v26a3 3 0 0 0 3 3h26a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3Z"/><path stroke-linecap="round" d="M44 13.002V42a2 2 0 0 1-2 2H13.003M13 20.486l6 5.525l10-10.292"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFullSelection0)"/>`),
		g.Group(children),
	)
}