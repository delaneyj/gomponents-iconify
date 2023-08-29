package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 .214l5 13l5-13L23.456 17H18v5h-2v-5H8v5H6v-5H.544L7 .214ZM8 15h2.544L8 8.385V15ZM6 8.385L3.456 15H6V8.385ZM13.456 15H16V8.385L13.456 15ZM18 8.385V15h2.544L18 8.385Z"/>`),
		g.Group(children),
	)
}