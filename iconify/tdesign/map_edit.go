package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V11h-2V5.5l-4 1.667V11h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926l7-4.084ZM8 16.833V4.741L4 7.074V18.5l4-1.667Zm11.787-4.747l4.127 4.127l-7.286 7.287H12.5l-.001-4.128l7.287-7.286Zm-.922 3.75l1.299 1.3l.922-.923l-1.3-1.299l-.921.922Zm-.115 2.713l-1.3-1.299l-2.95 2.95v1.3h1.3l2.95-2.95Z"/>`),
		g.Group(children),
	)
}