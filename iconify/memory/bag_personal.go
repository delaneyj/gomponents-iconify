package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagPersonal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 15H9v2H7v-2H5v4h12v-4m0-6h-1V8h-1V7H7v1H6v1H5v4h12V9m-4 2H9v-1h1V9h2v1h1v1M3 8h1V6h2V5h1V2h1V1h6v1h1v3h1v1h2v2h1v12h-1v1H4v-1H3V8m6-5v2h4V3H9Z"/>`),
		g.Group(children),
	)
}