package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChalkboardSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M28 168V56a12 12 0 0 1 12-12h176a12 12 0 0 1 12 12v112a4 4 0 0 1-8 0V56a4 4 0 0 0-4-4H40a4 4 0 0 0-4 4v112a4 4 0 0 1-8 0Zm216 32a4 4 0 0 1-4 4H16a4 4 0 0 1 0-8h100v-28a4 4 0 0 1 4-4h64a4 4 0 0 1 4 4v28h52a4 4 0 0 1 4 4Zm-120-4h56v-24h-56Z"/>`),
		g.Group(children),
	)
}