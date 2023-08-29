package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectCheckBoxRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.2 0 .375.038t.35.112L17.875 5H5v14h14v-6.65l2-2V19q0 .825-.588 1.413T19 21H5Zm6.525-6.8l8.5-8.5q.275-.275.675-.275t.7.275q.3.275.3.7t-.3.725L12.225 16.3q-.3.3-.7.3t-.7-.3l-4.25-4.25q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l3.55 3.55Z"/>`),
		g.Group(children),
	)
}