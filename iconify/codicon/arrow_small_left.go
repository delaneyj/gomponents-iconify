package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSmallLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.5 10.7L4 8.2v-.7L6.5 5l.71.7l-1.64 1.65h5.57v1H5.57L7.22 10l-.72.7z"/>`),
		g.Group(children),
	)
}