package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadgesOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.505 14.497L12 16l-5-3v4l5 3l5-3m-3.127-7.124L17 8V4l-5 3l-2.492-1.495M7 7v1l2.492 1.495M3 3l18 18"/>`),
		g.Group(children),
	)
}