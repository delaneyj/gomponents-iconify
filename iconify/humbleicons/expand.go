package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8.5V4m0 0h4.5M4 4l5.5 5.5m10.5-1V4m0 0h-4.5M20 4l-5.5 5.5M4 15.5V20m0 0h4.5M4 20l5.5-5.5m10.5 1V20m0 0h-4.5m4.5 0l-5.5-5.5"/>`),
		g.Group(children),
	)
}