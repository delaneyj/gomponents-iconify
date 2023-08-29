package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficSignalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 58a30 30 0 1 0 30 30a30 30 0 0 0-30-30Zm0 48a18 18 0 1 1 18-18a18 18 0 0 1-18 18Zm0 32a30 30 0 1 0 30 30a30 30 0 0 0-30-30Zm0 48a18 18 0 1 1 18-18a18 18 0 0 1-18 18Zm88-40h-18V78h18a6 6 0 0 0 0-12h-18V40a14 14 0 0 0-14-14H72a14 14 0 0 0-14 14v26H40a6 6 0 0 0 0 12h18v68H40a6 6 0 0 0 0 12h18v58a14 14 0 0 0 14 14h112a14 14 0 0 0 14-14v-58h18a6 6 0 0 0 0-12Zm-30 70a2 2 0 0 1-2 2H72a2 2 0 0 1-2-2V40a2 2 0 0 1 2-2h112a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}