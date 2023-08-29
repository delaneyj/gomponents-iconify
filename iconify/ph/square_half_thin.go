package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareHalfThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 44H56a12 12 0 0 0-12 12v144a12 12 0 0 0 12 12h144a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12Zm-68 72h72v24h-72Zm0-8V84h72v24Zm0 40h72v24h-72Zm72-92v20h-72V52h68a4 4 0 0 1 4 4ZM52 200V56a4 4 0 0 1 4-4h68v152H56a4 4 0 0 1-4-4Zm148 4h-68v-24h72v20a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}