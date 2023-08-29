package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M188.8 169.6L133.33 128l23.47-17.6a8 8 0 0 0-9.6-12.8L128 112V48l19.2 14.4a8 8 0 1 0 9.6-12.8l-32-24A8 8 0 0 0 112 32v80L60.8 73.6a8 8 0 0 0-9.6 12.8l55.47 41.6l-55.47 41.6a8 8 0 1 0 9.6 12.8L112 144v80a8 8 0 0 0 12.8 6.4l64-48a8 8 0 0 0 0-12.8ZM128 208v-64l42.67 32ZM237.66 98.34a8 8 0 0 1-11.32 11.32L208 91.31l-18.34 18.35a8 8 0 0 1-11.32-11.32L196.69 80l-18.35-18.34a8 8 0 0 1 11.32-11.32L208 68.69l18.34-18.35a8 8 0 0 1 11.32 11.32L219.31 80Z"/>`),
		g.Group(children),
	)
}