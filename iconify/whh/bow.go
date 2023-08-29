package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M895.998 768q14 0 28 14l100 100v14h-96q-13 0-22.5 9.5t-9.5 22.5v96h-14l-100-100q-14-14-14-28v-84l-256-256l-454 453q-9 15-26 15q-13 0-22.5-9.5t-9.5-22.5v-32q0-187 68.5-355.5t191.5-300.5l-186-186l-42 42q-13 0-22.5-9.5t-9.5-22.5V32q0-13 9.5-22.5t22.5-9.5h96q13 0 22.5 9.5t9.5 22.5l-43 43l187 186q131-124 300-192.5t356-68.5h32q13 0 22.5 9.5t9.5 22.5q0 17-15 26l-454 454l256 256h85zm-829 146l401-402l-121-121q-117 122-192.5 259.5T66.998 914zm847-847q-126 12-263.5 87.5T389.998 347l122 122z"/>`),
		g.Group(children),
	)
}