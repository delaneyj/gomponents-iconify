package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200.73 40H55.27A15.29 15.29 0 0 0 40 55.27v145.46A15.29 15.29 0 0 0 55.27 216h145.46A15.29 15.29 0 0 0 216 200.73V55.27A15.29 15.29 0 0 0 200.73 40ZM200 200H56V56h144Z"/>`),
		g.Group(children),
	)
}