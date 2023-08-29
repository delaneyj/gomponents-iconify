package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessFiveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.675 20h-2.65q-.825 0-1.413-.588T4.026 18v-2.65L2.1 13.4q-.575-.575-.575-1.4t.575-1.4l1.925-1.95V6q0-.825.588-1.413T6.025 4h2.65l1.95-1.925q.575-.575 1.4-.575t1.4.575L15.375 4h2.65q.825 0 1.413.588T20.024 6v2.65l1.925 1.95q.575.575.575 1.4t-.575 1.4l-1.925 1.95V18q0 .825-.587 1.413T18.025 20h-2.65l-1.95 1.925q-.575.575-1.4.575t-1.4-.575L8.675 20Zm3.35-3q-2.075 0-3.538-1.463T7.026 12q0-2.075 1.463-3.538T12.024 7q2.075 0 3.538 1.463T17.024 12q0 2.075-1.462 3.538T12.025 17Zm0 3.5l2.5-2.5h3.5v-3.5l2.5-2.5l-2.5-2.5V6h-3.5l-2.5-2.5l-2.5 2.5h-3.5v3.5l-2.5 2.5l2.5 2.5V18h3.5l2.5 2.5Z"/>`),
		g.Group(children),
	)
}