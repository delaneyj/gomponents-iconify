package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiNoneThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M136 204a8 8 0 1 1-8-8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}