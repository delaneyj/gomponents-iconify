package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditSquareRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.875 9.05L17.6 4.85l1.075-1.075q.6-.6 1.438-.6t1.412.6l1.4 1.425q.575.575.575 1.4T22.925 8l-1.05 1.05ZM5 23.7q-.825 0-1.413-.587T3 21.7v-14q0-.825.588-1.413T5 5.7h8.925l-6.05 6.05q-.425.425-.65.975T7 13.875V17.7q0 .825.588 1.413T9 19.7h3.825q.6 0 1.15-.225t.975-.65L21 12.75v8.95q0 .825-.588 1.413T19 23.7H5Zm5-6q-.425 0-.713-.287T9 16.7v-2.425q0-.4.15-.763t.425-.637l6.6-6.6l4.275 4.2l-6.625 6.625q-.275.275-.637.438t-.763.162H10Z"/>`),
		g.Group(children),
	)
}