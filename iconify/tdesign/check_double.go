package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.657 7.171l-5.303 5.304l-1.414-1.415l5.303-5.303l1.414 1.414Zm5.657 0L12.708 17.778l-6.364-6.364L7.758 10l4.95 4.95L21.9 5.757l1.414 1.414ZM2.101 10l4.95 4.95l.353-.354l1.414 1.414l-1.768 1.768l-6.363-6.364L2.1 10Z"/>`),
		g.Group(children),
	)
}