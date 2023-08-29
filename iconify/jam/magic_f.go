package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagicF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.092 8.02l-2.829-2.828L16.506.95a1 1 0 0 1 1.414 0l1.415 1.414a1 1 0 0 1 0 1.414l-4.243 4.243zm-1.414 1.415l-9.9 9.9a1 1 0 0 1-1.414 0L.95 17.92a1 1 0 0 1 0-1.414l9.9-9.9l2.828 2.829zM8.728.243l1.393.704l1.435-.704l-.679 1.46l.68 1.368l-1.384-.664l-1.445.664l.689-1.42l-.69-1.408zm9.9 7.07l1.393.705l1.435-.704l-.68 1.46l.68 1.368l-1.384-.664l-1.445.664l.69-1.42l-.69-1.408z"/>`),
		g.Group(children),
	)
}