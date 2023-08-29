package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotVisible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 12s4.188-6 11-6c.947 0 1.839.121 2.678.322L8.36 12.64A3.96 3.96 0 0 1 8 11c0-.937.335-1.787.875-2.469c-2.392.812-4.214 2.385-5.254 3.469a14.868 14.868 0 0 0 2.98 2.398l-1.453 1.453C2.497 14.13 1 12 1 12zm22 0s-4.188 6-11 6c-.946 0-1.836-.124-2.676-.323L5 22l-1.5-1.5l17-17L22 5l-3.147 3.147C21.501 9.869 23 12 23 12zm-2.615.006a14.83 14.83 0 0 0-2.987-2.403L16 11a4 4 0 0 1-4 4l-.947.947c.31.031.624.053.947.053c3.978 0 6.943-2.478 8.385-3.994z"/>`),
		g.Group(children),
	)
}