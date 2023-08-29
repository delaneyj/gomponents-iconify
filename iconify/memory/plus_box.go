package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 16h-2v-4H6v-2h4V6h2v4h4v2h-4Zm6 4H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-1-2v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		g.Group(children),
	)
}