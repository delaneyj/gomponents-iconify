package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChristmasTreeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 36V44"/><path fill="#2F88FF" d="M14 15L24 4L34 15L31 18L39 26L34 27.1579L42 36H6L14 27.1579L9 26L17 18L14 15Z"/></g>`),
		g.Group(children),
	)
}