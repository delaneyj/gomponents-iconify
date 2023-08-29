package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.606 2.08a1 1 0 0 1 .788 0l6.394 2.741A2 2 0 0 1 20 6.66v6.684a1 1 0 0 1-2 0V6.659l-6-2.571l-1.946.834a1 1 0 1 1-.788-1.838l2.34-1.003ZM4.26 5.675c-.167.294-.26.632-.26.985v6.86a7 7 0 0 0 3.527 6.077l3.977 2.272a1 1 0 0 0 .992 0l3.977-2.272c.36-.206.696-.44 1.009-.7l1.811 1.811a1 1 0 0 0 1.414-1.414L6.422 5.008a.98.98 0 0 0-.03-.03L4.707 3.293a1 1 0 0 0-1.414 1.414l.967.967Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}