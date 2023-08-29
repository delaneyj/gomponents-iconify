package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21H5q-.825 0-1.413-.588T3 19v-7q0-1.875.713-3.513t1.924-2.85q1.213-1.212 2.85-1.924T12 3q1.875 0 3.513.713t2.85 1.924q1.212 1.213 1.925 2.85T21 12v7q0 .825-.588 1.413T19 21h-4v-8h4v-1q0-2.925-2.038-4.963T12 5Q9.075 5 7.037 7.038T5 12v1h4v8Z"/>`),
		g.Group(children),
	)
}