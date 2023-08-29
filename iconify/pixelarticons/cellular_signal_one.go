package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellularSignalOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 4v16h6V4h-6Zm2 2h2v12h-2V6Zm-9 4v10h6V10H9Zm2 8v-6h2v6h-2Zm-3-4H2v6h6v-6Z"/>`),
		g.Group(children),
	)
}