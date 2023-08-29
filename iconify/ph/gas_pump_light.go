package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasPumpLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m239.56 71.07l-19.32-19.31a6 6 0 0 0-8.48 8.48l19.31 19.32a9.93 9.93 0 0 1 2.93 7.07V168a10 10 0 0 1-20 0v-40a22 22 0 0 0-22-22h-18V56a22 22 0 0 0-22-22H72a22 22 0 0 0-22 22v154H32a6 6 0 0 0 0 12h160a6 6 0 0 0 0-12h-18v-92h18a10 10 0 0 1 10 10v40a22 22 0 0 0 44 0V86.63a21.88 21.88 0 0 0-6.44-15.56ZM62 210V56a10 10 0 0 1 10-10h80a10 10 0 0 1 10 10v154Zm80-98a6 6 0 0 1-6 6H88a6 6 0 0 1 0-12h48a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}