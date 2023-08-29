package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpAirplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22h12l-6-6l-6 6zM23 3H1v16h6v-2H3V5h18v12h-4v2h6V3z"/>`),
		g.Group(children),
	)
}