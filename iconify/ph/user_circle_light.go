package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26ZM71.44 198a66 66 0 0 1 113.12 0a89.8 89.8 0 0 1-113.12 0ZM94 120a34 34 0 1 1 34 34a34 34 0 0 1-34-34Zm99.51 69.64a77.53 77.53 0 0 0-40-31.38a46 46 0 1 0-51 0a77.53 77.53 0 0 0-40 31.38a90 90 0 1 1 131 0Z"/>`),
		g.Group(children),
	)
}