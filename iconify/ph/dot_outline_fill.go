package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotOutlineFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M156 128a28 28 0 1 1-28-28a28 28 0 0 1 28 28Z"/>`),
		g.Group(children),
	)
}