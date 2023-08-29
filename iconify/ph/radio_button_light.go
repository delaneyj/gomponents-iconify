package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioButtonLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm0-152a62 62 0 1 0 62 62a62.07 62.07 0 0 0-62-62Zm0 112a50 50 0 1 1 50-50a50.06 50.06 0 0 1-50 50Z"/>`),
		g.Group(children),
	)
}