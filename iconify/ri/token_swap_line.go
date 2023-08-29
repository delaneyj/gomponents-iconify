package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TokenSwapLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 12.5L6.5 15L9 17.5l2.5-2.5L9 12.5Zm6-10a6.5 6.5 0 0 0-6.482 6.018a6.5 6.5 0 1 0 6.964 6.964A6.5 6.5 0 0 0 15 2.5Zm.323 10.989a6.51 6.51 0 0 0-4.812-4.812a4.5 4.5 0 1 1 4.812 4.812ZM13.5 15a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0ZM3 7a4 4 0 0 1 4-4h1.5v2H7a2 2 0 0 0-2 2v1.5H3V7Zm16 10v-1.5h2V17a4 4 0 0 1-4 4h-1.5v-2H17a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}