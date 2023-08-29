package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Zm40-72a8 8 0 0 1-8 8h-8v24a8 8 0 0 1-16 0v-24H96a8 8 0 0 1-7.59-10.53l24-72a8 8 0 0 1 15.18 5.06L107.1 136H136v-24a8 8 0 0 1 16 0v24h8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}