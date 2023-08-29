package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.29 20.69a.78.78 0 0 0 .24-.16l4-4a.75.75 0 0 0-1.06-1.06l-2.72 2.72V5.81l2.72 2.72a.75.75 0 0 0 1.06-1.06l-4-4a.78.78 0 0 0-.24-.16a.73.73 0 0 0-.58 0a.78.78 0 0 0-.24.16l-4 4a.75.75 0 0 0 0 1.06a.75.75 0 0 0 1.06 0l2.72-2.72v12.38l-2.72-2.72a.75.75 0 0 0-1.06 0a.75.75 0 0 0 0 1.06l4 4a.78.78 0 0 0 .24.16a.73.73 0 0 0 .58 0Z"/>`),
		g.Group(children),
	)
}