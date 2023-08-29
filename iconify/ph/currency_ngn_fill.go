package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyNgnFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 93.63L110.61 120H88Zm80 68.74V136h-22.61ZM232 128A104 104 0 1 1 128 24a104.11 104.11 0 0 1 104 104Zm-32 0a8 8 0 0 0-8-8h-8V72a8 8 0 0 0-16 0v48h-36.32L86.07 66.79A8 8 0 0 0 72 72v48h-8a8 8 0 0 0 0 16h8v48a8 8 0 0 0 16 0v-48h36.32l45.61 53.21A8 8 0 0 0 176 192a8 8 0 0 0 8-8v-48h8a8 8 0 0 0 8-8Z"/>`),
		g.Group(children),
	)
}