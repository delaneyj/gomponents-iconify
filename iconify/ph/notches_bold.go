package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotchesBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m216.49 136.49l-80 80a12 12 0 1 1-17-17l80-80a12 12 0 1 1 17 17Zm-16-105a12 12 0 0 0-17 0l-152 152a12 12 0 0 0 17 17l152-152a12 12 0 0 0 0-16.98Z"/>`),
		g.Group(children),
	)
}