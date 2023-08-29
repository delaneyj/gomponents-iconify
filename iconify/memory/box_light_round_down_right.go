package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightRoundDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 0v4h-1v3h-1v1H9v1H8v1H7v1H4v1H0v-2h3V9h3V8h1V7h1V6h1V3h1V0h2Z"/>`),
		g.Group(children),
	)
}