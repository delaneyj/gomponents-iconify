package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowDownRightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m208 176l-48 48v-96Z" opacity=".2"/><path d="m213.66 170.34l-48-48A8 8 0 0 0 152 128v40H72V32a8 8 0 0 0-16 0v144a8 8 0 0 0 8 8h88v40a8 8 0 0 0 13.66 5.66l48-48a8 8 0 0 0 0-11.32ZM168 204.69v-57.38L196.69 176Z"/></g>`),
		g.Group(children),
	)
}