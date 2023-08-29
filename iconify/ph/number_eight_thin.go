package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M147.08 119.64a44 44 0 1 0-38.16 0a52 52 0 1 0 38.16 0ZM92 80a36 36 0 1 1 36 36a36 36 0 0 1-36-36Zm36 132a44 44 0 1 1 44-44a44.05 44.05 0 0 1-44 44Z"/>`),
		g.Group(children),
	)
}