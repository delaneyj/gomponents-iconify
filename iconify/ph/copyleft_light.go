package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyleftLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm46-90a46 46 0 0 1-82.8 27.61a6 6 0 0 1 9.6-7.21a34 34 0 1 0 0-40.8a6 6 0 0 1-9.6-7.21A46 46 0 0 1 174 128Z"/>`),
		g.Group(children),
	)
}