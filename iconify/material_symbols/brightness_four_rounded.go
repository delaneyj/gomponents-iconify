package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessFourRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.3 17q2.075 0 3.538-1.463T17.3 12q0-2.075-1.463-3.538T12.3 7h-.413q-.212 0-.412.05q-.375.075-.475.438t.175.562q.95.7 1.45 1.738t.5 2.212q0 1.2-.475 2.263t-1.475 1.687q-.3.2-.213.563t.513.412q.2.05.413.063T12.3 17Zm3.05 3l-1.95 1.925q-.575.575-1.4.575t-1.4-.575L8.65 20H6q-.825 0-1.413-.588T4 18v-2.65L2.075 13.4Q1.5 12.825 1.5 12t.575-1.4L4 8.65V6q0-.825.588-1.413T6 4h2.65l1.95-1.925Q11.175 1.5 12 1.5t1.4.575L15.35 4H18q.825 0 1.413.588T20 6v2.65l1.925 1.95q.575.575.575 1.4t-.575 1.4L20 15.35V18q0 .825-.588 1.413T18 20h-2.65Z"/>`),
		g.Group(children),
	)
}