package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessAutoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.925 7.75l-2.75 7.225q-.15.35.075.688t.625.337q.25 0 .438-.138t.262-.362l.625-1.8h3.65l.625 1.825q.075.225.263.35t.412.125q.375 0 .588-.313t.087-.662l-2.75-7.275q-.125-.35-.425-.55T12 7q-.35 0-.65.2t-.425.55Zm-.275 4.65l1.3-3.75h.1l1.3 3.75h-2.7Zm-2 7.6H6q-.825 0-1.413-.588T4 18v-2.65L2.075 13.4Q1.5 12.825 1.5 12t.575-1.4L4 8.65V6q0-.825.588-1.413T6 4h2.65l1.95-1.925Q11.175 1.5 12 1.5t1.4.575L15.35 4H18q.825 0 1.413.588T20 6v2.65l1.925 1.95q.575.575.575 1.4t-.575 1.4L20 15.35V18q0 .825-.588 1.413T18 20h-2.65l-1.95 1.925q-.575.575-1.4.575t-1.4-.575L8.65 20Z"/>`),
		g.Group(children),
	)
}