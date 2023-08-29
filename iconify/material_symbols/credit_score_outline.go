package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditScoreOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 8h16V6H4v2ZM2 6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v6H4v6h4.1v2H4q-.825 0-1.413-.588T2 18V6Zm12.95 16l-4.25-4.25l1.4-1.4l2.85 2.8l5.65-5.65l1.4 1.45L14.95 22ZM4 6v12v-4.5v2.825V6Z"/>`),
		g.Group(children),
	)
}