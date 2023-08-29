package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThinUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 3.828L2.929 9.899L1.515 8.485L10 0l.707.707l7.778 7.778l-1.414 1.414L11 3.828V20H9V3.828z"/>`),
		g.Group(children),
	)
}