package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 3h1V2h13v1h1v1h1v1h1v1h1v13h-1v1H3v-1H2V3m16 4h-1V6h-1V5h-1v4H6V4H4v14h2v-5h10v5h2V7m-7-3v3h2V4h-2m3 14v-3H8v3h6Z"/>`),
		g.Group(children),
	)
}