package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartHalfDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M128 72v144S24 160 24 94a54 54 0 0 1 54-54c22.59 0 41.94 12.31 50 32Z" opacity=".2"/><path d="M178 32c-20.65 0-38.73 8.88-50 23.89C116.73 40.88 98.65 32 78 32a62.07 62.07 0 0 0-62 62c0 70 103.79 126.67 108.21 129a8 8 0 0 0 7.58 0C136.21 220.67 240 164 240 94a62.07 62.07 0 0 0-62-62ZM32 94a46.06 46.06 0 0 1 46-46c18.91 0 34.86 9.79 42 25.65V202c-26.41-16.56-88-60.22-88-108Zm104 108V73.65C143.14 57.79 159.09 48 178 48a46.06 46.06 0 0 1 46 46c0 47.71-61.58 91.41-88 108Z"/></g>`),
		g.Group(children),
	)
}