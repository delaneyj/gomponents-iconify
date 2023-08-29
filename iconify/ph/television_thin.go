package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 68h-78.34l41.17-41.17a4 4 0 1 0-5.66-5.66L128 66.34L82.83 21.17a4 4 0 0 0-5.66 5.66L118.34 68H40a12 12 0 0 0-12 12v120a12 12 0 0 0 12 12h176a12 12 0 0 0 12-12V80a12 12 0 0 0-12-12ZM36 200V80a4 4 0 0 1 4-4h108v128H40a4 4 0 0 1-4-4Zm184 0a4 4 0 0 1-4 4h-60V76h60a4 4 0 0 1 4 4Zm-24-84a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm0 48a8 8 0 1 1-8-8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}