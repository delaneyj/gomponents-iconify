package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightDownLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m228.24 164.24l-48 48a6 6 0 0 1-8.48 0l-48-48a6 6 0 1 1 8.48-8.48L170 193.51V70H32a6 6 0 0 1 0-12h144a6 6 0 0 1 6 6v129.51l37.76-37.75a6 6 0 0 1 8.48 8.48Z"/>`),
		g.Group(children),
	)
}