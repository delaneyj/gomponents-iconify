package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudControllerManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 10h8v2H9Zm15 0h-3V8h-2v6h2v-2h3ZM9 20h8v2H9Zm15 0h-3v-2h-2v6h2v-2h3Zm0-3h-8v-2h8ZM9 17h3v2h2v-6h-2v2H9Z"/><path fill="currentColor" d="M19 6H7v10s-7 .17-7-6a6.071 6.071 0 0 1 1.235-4a5.99 5.99 0 0 1 4.459-1.99a4.86 4.86 0 0 1 .572-.99l.272-.36l.23-.26a3.597 3.597 0 0 1 .262-.29A6.897 6.897 0 0 1 12 0c.16 0 .362.01.513.02a7.012 7.012 0 0 1 3.756 1.39a8.546 8.546 0 0 1 .913.81A7.043 7.043 0 0 1 19 6Z"/>`),
		g.Group(children),
	)
}