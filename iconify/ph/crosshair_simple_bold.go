package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrosshairSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm12 191.13V184a12 12 0 0 0-24 0v27.13A84.18 84.18 0 0 1 44.87 140H72a12 12 0 0 0 0-24H44.87A84.18 84.18 0 0 1 116 44.87V72a12 12 0 0 0 24 0V44.87A84.18 84.18 0 0 1 211.13 116H184a12 12 0 0 0 0 24h27.13A84.18 84.18 0 0 1 140 211.13Z"/>`),
		g.Group(children),
	)
}