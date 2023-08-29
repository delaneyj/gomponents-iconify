package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 13a7 7 0 1 0 14 0a7 7 0 1 0-14 0"/><path d="M12 10v3h2M7 4L4.25 6M17 4l2.75 2"/></g>`),
		g.Group(children),
	)
}