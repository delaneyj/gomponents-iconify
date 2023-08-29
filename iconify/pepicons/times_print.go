package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimesPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.132 6.477a1.5 1.5 0 0 0 .073 2.12L8.78 11l-2.575 2.403a1.5 1.5 0 0 0 2.047 2.194l2.727-2.545l2.726 2.545a1.5 1.5 0 0 0 2.047-2.194L13.177 11l2.575-2.403a1.5 1.5 0 1 0-2.047-2.194L10.98 8.948L8.252 6.403a1.5 1.5 0 0 0-2.12.074Z" clip-rule="evenodd" opacity=".8"/><path d="M6.854 13.854a.5.5 0 0 1-.708-.708l7-7a.5.5 0 0 1 .708.708l-7 7Z"/><path d="M6.146 6.854a.5.5 0 1 1 .708-.708l7 7a.5.5 0 0 1-.708.708l-7-7Z"/></g>`),
		g.Group(children),
	)
}