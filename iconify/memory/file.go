package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func File(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M13 1v1h1v1h1v1h1v1h1v1h1v1h1v13h-1v1H4v-1H3V2h1V1h9m0 3h-1v4h4V7h-1V6h-1V5h-1V4M5 3v16h12v-9h-6V9h-1V3H5Z"/>`),
		g.Group(children),
	)
}