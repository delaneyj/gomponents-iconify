package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardholderThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 52H48a20 20 0 0 0-20 20v112a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V72a20 20 0 0 0-20-20ZM36 92h184v24h-52.81a12 12 0 0 0-11.75 9.6a28 28 0 0 1-54.88 0a12 12 0 0 0-11.75-9.6H36Zm12-32h160a12 12 0 0 1 12 12v12H36V72a12 12 0 0 1 12-12Zm160 136H48a12 12 0 0 1-12-12v-60h52.81a4 4 0 0 1 3.91 3.2a36 36 0 0 0 70.56 0a4 4 0 0 1 3.91-3.2H220v60a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}