package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v8h-2V4H4v9.586l5-5l5.914 5.914l-1.414 1.414l-4.5-4.5l-5 5V20h6v2H2V2Zm13.547 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm7.24 4.086l4.127 4.127l-7.286 7.287H12.5l-.001-4.128l7.287-7.286Zm-.922 3.75l1.299 1.3l.922-.923l-1.3-1.299l-.921.922Zm-.115 2.713l-1.3-1.299l-2.95 2.95v1.3h1.3l2.95-2.95Z"/>`),
		g.Group(children),
	)
}