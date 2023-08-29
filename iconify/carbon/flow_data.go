package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 23h-8.14a4.17 4.17 0 0 0-.43-1L22 11.43a3.86 3.86 0 0 0 2 .57a4 4 0 1 0-3.86-5h-8.28a4 4 0 1 0 0 2h8.28a4.17 4.17 0 0 0 .43 1L10 20.57A3.86 3.86 0 0 0 8 20a4 4 0 1 0 3.86 5H20v3h8v-8h-8ZM8 10a2 2 0 1 1 2-2a2 2 0 0 1-2 2Zm16-4a2 2 0 1 1-2 2a2 2 0 0 1 2-2ZM8 26a2 2 0 1 1 2-2a2 2 0 0 1-2 2Zm14-4h4v4h-4Z"/>`),
		g.Group(children),
	)
}