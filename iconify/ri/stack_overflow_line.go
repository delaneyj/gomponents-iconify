package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackOverflowLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.001 20.003V15h2v7.003h-16V15h2v5.003h12ZM7.501 18v-2h9v2h-9Zm.077-4.38l.347-1.97l8.864 1.563l-.348 1.97l-8.863-1.563Zm1.634-5.504l1-1.732l7.794 4.5l-1 1.732l-7.794-4.5Zm3.417-4.613l1.532-1.285l5.785 6.894l-1.532 1.286l-5.785-6.895Z"/>`),
		g.Group(children),
	)
}