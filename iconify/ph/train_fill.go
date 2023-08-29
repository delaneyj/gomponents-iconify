package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 24H72a32 32 0 0 0-32 32v128a32 32 0 0 0 32 32h8l-14.4 19.2a8 8 0 1 0 12.8 9.6L100 216h56l21.6 28.8a8 8 0 1 0 12.8-9.6L176 216h8a32 32 0 0 0 32-32V56a32 32 0 0 0-32-32ZM84 184a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm36-64H56V80h64Zm52 64a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm28-64h-64V80h64Z"/>`),
		g.Group(children),
	)
}