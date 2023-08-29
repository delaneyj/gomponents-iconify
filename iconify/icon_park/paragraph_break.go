package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 4V44"/><path d="M42 4V44"/><path d="M18 26L14 30L18 34"/><path d="M15 30C15 30 25.7909 30 28 30C31.3137 30 34 27.3137 34 24C34 20.6863 31.3137 18 28 18C26.6852 18 14 18 14 18"/></g>`),
		g.Group(children),
	)
}