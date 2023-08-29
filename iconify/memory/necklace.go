package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Necklace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M9 17h1v-1h2v1h1v2h-1v1h-2v-1H9v-2m1-2v-1H9v-1h-.91v-1H7v-2H6V8H5V6H4V3h1V2h12v1h1v3h-1v2h-1v2h-1v2h-1v1h-1v1h-1v1h-2M7 5v2h1v2h1v2h1.09v1H12v-1h1V9h1V7h1V5h1V4H6v1h1Z"/>`),
		g.Group(children),
	)
}