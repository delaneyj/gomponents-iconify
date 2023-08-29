package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesThreePlusFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 80a40 40 0 1 1-40-40a40 40 0 0 1 40 40Zm56 40a40 40 0 1 0-40-40a40 40 0 0 0 40 40Zm-96 16a40 40 0 1 0 40 40a40 40 0 0 0-40-40Zm128 32h-24v-24a8 8 0 0 0-16 0v24h-24a8 8 0 0 0 0 16h24v24a8 8 0 0 0 16 0v-24h24a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}