package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.69 11.71a.78.78 0 0 0-.16-.24l-4-4a.75.75 0 1 0-1.06 1.06l2.72 2.72H5.81l2.72-2.72a.75.75 0 0 0-1.06-1.06l-4 4a.78.78 0 0 0-.16.24a.73.73 0 0 0 0 .58a.78.78 0 0 0 .16.24l4 4a.75.75 0 0 0 1.06 0a.75.75 0 0 0 0-1.06l-2.72-2.72h12.38l-2.72 2.72a.75.75 0 0 0 0 1.06a.75.75 0 0 0 1.06 0l4-4a.78.78 0 0 0 .16-.24a.73.73 0 0 0 0-.58Z"/>`),
		g.Group(children),
	)
}