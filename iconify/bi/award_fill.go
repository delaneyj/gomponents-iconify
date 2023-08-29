package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AwardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="m8 0l1.669.864l1.858.282l.842 1.68l1.337 1.32L13.4 6l.306 1.854l-1.337 1.32l-.842 1.68l-1.858.282L8 12l-1.669-.864l-1.858-.282l-.842-1.68l-1.337-1.32L2.6 6l-.306-1.854l1.337-1.32l.842-1.68L6.331.864L8 0z"/><path d="M4 11.794V16l4-1l4 1v-4.206l-2.018.306L8 13.126L6.018 12.1L4 11.794z"/></g>`),
		g.Group(children),
	)
}