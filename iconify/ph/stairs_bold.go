package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StairsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 20H56a20 20 0 0 0-20 20v176a20 20 0 0 0 20 20h144a20 20 0 0 0 20-20V40a20 20 0 0 0-20-20Zm-48 128h44v16h-80v-16Zm12-24v-16h32v16Zm32-80v40h-44a12 12 0 0 0-12 12v28h-36a12 12 0 0 0-12 12v28H60V44ZM60 212v-24h136v24Z"/>`),
		g.Group(children),
	)
}