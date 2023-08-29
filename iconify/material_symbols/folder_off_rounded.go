package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.775 18.95L6.85 4h2.025q.4 0 .763.15t.637.425L11.7 6H20q.825 0 1.413.588T22 8v10q0 .275-.05.513t-.175.437ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4l2 2H3.2L1.4 4.2q-.275-.275-.287-.688T1.4 2.8q.275-.275.7-.275t.7.275l18.4 18.4q.3.3.288.7t-.313.7q-.3.275-.7.288t-.7-.288l-2.6-2.6H4Z"/>`),
		g.Group(children),
	)
}