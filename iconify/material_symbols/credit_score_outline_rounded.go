package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditScoreOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 8h16V6H4v2ZM2 6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v6H4v6h3.1q.425 0 .713.288T8.1 19q0 .425-.288.713T7.1 20H4q-.825 0-1.413-.588T2 18V6Zm2 0v12v-3.263v1.588V6Zm10.95 13.15l4.925-4.925q.3-.3.713-.288t.712.313q.275.3.288.7t-.288.7l-5.65 5.65q-.3.3-.7.3t-.7-.3l-2.85-2.85q-.275-.275-.288-.688t.288-.712q.275-.275.688-.275t.712.275l2.15 2.1Z"/>`),
		g.Group(children),
	)
}