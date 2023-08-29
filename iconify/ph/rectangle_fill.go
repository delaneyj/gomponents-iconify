package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<rect width="208" height="176" x="24" y="40" fill="currentColor" rx="16"/>`),
		g.Group(children),
	)
}