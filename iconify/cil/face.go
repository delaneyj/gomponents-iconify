package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Face(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16Zm147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125Z"/><path fill="currentColor" d="M128 192h40v40h-40zm208 0h40v40h-40zM232 306.948a5.049 5.049 0 0 1 .194-1.387L264 194.241V168h-32v21.759l-30.574 107.009A37.052 37.052 0 0 0 237.052 344H296v-32h-58.948a5.057 5.057 0 0 1-5.052-5.052Z"/>`),
		g.Group(children),
	)
}