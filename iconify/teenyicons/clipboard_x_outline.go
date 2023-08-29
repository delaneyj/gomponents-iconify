package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardXOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M11 1.5h2.5v12a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1v-12H4m1.5 5l4 4m-4 0l4-4m-5-6h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2Z"/>`),
		g.Group(children),
	)
}