package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanDeleteOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4v5v-5v16v-.238V20V4Zm0 18q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762v3.525q-.475-.175-.975-.263T18 12V9h-3q-.825 0-1.413-.588T13 7V4H6v16h6.35q.2.575.5 1.075t.7.925H6Zm12-2.575L16.6 20.8q-.275.275-.688.288T15.2 20.8q-.275-.275-.275-.7t.275-.7l1.4-1.4l-1.4-1.4q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.4 1.4l1.4-1.4q.275-.275.688-.287t.712.287q.275.275.275.7t-.275.7L19.425 18l1.375 1.4q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275L18 19.425Z"/>`),
		g.Group(children),
	)
}