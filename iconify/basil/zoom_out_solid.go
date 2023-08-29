package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.399 14.945a6.751 6.751 0 0 0 8.986.5l5.156 5.156a.75.75 0 0 0 1.06-1.06l-5.155-5.156a6.75 6.75 0 1 0-10.047.56ZM7.7 9.45a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5h-5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}