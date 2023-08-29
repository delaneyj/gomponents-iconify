package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseRollingFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 48h-24V24a24 24 0 0 0-24-24h-32a24 24 0 0 0-24 24v24H64a16 16 0 0 0-16 16v144a16 16 0 0 0 16 16h16v16a8 8 0 0 0 16 0v-16h64v16a8 8 0 0 0 16 0v-16h16a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16ZM96 192a8 8 0 0 1-16 0V80a8 8 0 0 1 16 0Zm40 0a8 8 0 0 1-16 0V80a8 8 0 0 1 16 0Zm16-144h-48V24a8 8 0 0 1 8-8h32a8 8 0 0 1 8 8Zm24 144a8 8 0 0 1-16 0V80a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}