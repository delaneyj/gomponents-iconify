package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalfLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm6 12.2a89.86 89.86 0 0 1 20 3.63v172.34a89.86 89.86 0 0 1-20 3.63Zm32 8.23a90.48 90.48 0 0 1 20 12.81v137.52a90.48 90.48 0 0 1-20 12.81ZM38 128a90.12 90.12 0 0 1 84-89.8v179.6A90.12 90.12 0 0 1 38 128Zm160 56.5v-113a89.81 89.81 0 0 1 0 113Z"/>`),
		g.Group(children),
	)
}