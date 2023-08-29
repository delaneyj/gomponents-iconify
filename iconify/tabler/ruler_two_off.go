package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerTwoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.03 7.97L17 3l4 4l-5 5m-2 2l-7 7l-4-4l7-7m6-3l-1.5-1.5M10 13l-1.5-1.5M7 16l-1.5-1.5M3 3l18 18"/>`),
		g.Group(children),
	)
}