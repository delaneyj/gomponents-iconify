package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowStream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 24h-8.137a4.13 4.13 0 0 0-.429-1.02L21.98 12.434A3.994 3.994 0 1 0 20.141 8H14v2h6.141a3.96 3.96 0 0 0 .425 1.02L10.02 21.566A3.994 3.994 0 1 0 11.859 26H20v3h8v-8h-8zm4-17a2 2 0 1 1-2 2a2.002 2.002 0 0 1 2-2zM8 27a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2zm14-4h4v4h-4zM9.693 12.75a5 5 0 0 1 0-7.5l1.324 1.5a3 3 0 0 0 0 4.501z"/><path fill="currentColor" d="M7.047 15.751a9 9 0 0 1 0-13.501l1.324 1.5a7 7 0 0 0 0 10.501Z"/>`),
		g.Group(children),
	)
}