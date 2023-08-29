package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 4H8V44H40V4Z"/><path d="M24 12V36"/><path d="M40 36H24H8"/><path d="M8 4L14 12H34L40 4"/></g>`),
		g.Group(children),
	)
}