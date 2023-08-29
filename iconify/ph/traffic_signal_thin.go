package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficSignalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 60a28 28 0 1 0 28 28a28 28 0 0 0-28-28Zm0 48a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm0 32a28 28 0 1 0 28 28a28 28 0 0 0-28-28Zm0 48a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm88-40h-20V76h20a4 4 0 0 0 0-8h-20V40a12 12 0 0 0-12-12H72a12 12 0 0 0-12 12v28H40a4 4 0 0 0 0 8h20v72H40a4 4 0 0 0 0 8h20v60a12 12 0 0 0 12 12h112a12 12 0 0 0 12-12v-60h20a4 4 0 0 0 0-8Zm-28 68a4 4 0 0 1-4 4H72a4 4 0 0 1-4-4V40a4 4 0 0 1 4-4h112a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}