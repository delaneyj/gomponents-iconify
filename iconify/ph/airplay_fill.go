package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplayFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M182.07 210.79A8 8 0 0 1 176 224H80a8 8 0 0 1-6.07-13.21l48-56a8 8 0 0 1 12.15 0ZM208 40H48a24 24 0 0 0-24 24v112a24 24 0 0 0 24 24h12.26a4 4 0 0 0 3-1.4l46.48-54.22a24 24 0 0 1 36.44 0l46.52 54.22a4 4 0 0 0 3 1.4H208a24 24 0 0 0 24-24V64a24 24 0 0 0-24-24Z"/>`),
		g.Group(children),
	)
}