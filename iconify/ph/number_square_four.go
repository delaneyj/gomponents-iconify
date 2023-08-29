package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSquareFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm0 176H48V48h160v160Zm-40-64a8 8 0 0 1-8 8h-8v24a8 8 0 0 1-16 0v-24H96a8 8 0 0 1-7.59-10.53l24-72a8 8 0 0 1 15.18 5.06L107.1 136H136v-24a8 8 0 0 1 16 0v24h8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}