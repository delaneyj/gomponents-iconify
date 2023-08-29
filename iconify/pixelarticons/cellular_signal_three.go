package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellularSignalThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 4h6v16h-6V4ZM2 14h6v6H2v-6Zm13-4H9v10h6V10Z"/>`),
		g.Group(children),
	)
}