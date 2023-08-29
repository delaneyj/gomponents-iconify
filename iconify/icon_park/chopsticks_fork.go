package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChopsticksFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M14 4V44"/><path d="M8 5V15C8 20 14 20 14 20C14 20 20 20 20 15V5"/><path d="M37 4L40 44"/><path d="M31 4L28 44"/></g>`),
		g.Group(children),
	)
}