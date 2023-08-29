package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraVideoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M3 5h14v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5Zm14 4l2.758-.69A1 1 0 0 1 21 9.28v5.44a1 1 0 0 1-1.242.97L17 15V9Z"/>`),
		g.Group(children),
	)
}