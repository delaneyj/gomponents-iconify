package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagPersonalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 8h1V6h2V5h1V2h1V1h6v1h1v3h1v1h2v2h1v12h-1v1H4v-1H3V8m6-5v2h4V3H9m8 11H5v1h2v2h1v-2h9v-1m-5-3h1V9h-1V8h-2v1H9v2h1v1h2v-1Z"/>`),
		g.Group(children),
	)
}