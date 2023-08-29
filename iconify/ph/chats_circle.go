package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatsCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M231.79 187.33a80 80 0 0 0-62.22-114.74a80 80 0 1 0-145.36 66.74l-7.66 26.82a14 14 0 0 0 17.3 17.3l26.82-7.66a80.15 80.15 0 0 0 25.75 7.63a80 80 0 0 0 108.91 40.37l26.82 7.66a14 14 0 0 0 17.3-17.3Zm-170.26-28.1a8.22 8.22 0 0 0-2.2.3l-26.41 7.55l7.55-26.41a8 8 0 0 0-.68-6a63.95 63.95 0 1 1 25.57 25.57a7.94 7.94 0 0 0-3.83-1.01Zm154 29.44l7.55 26.41l-26.41-7.55a8 8 0 0 0-6 .68a64.06 64.06 0 0 1-86.32-24.64a79.93 79.93 0 0 0 70.35-93.86a64 64 0 0 1 41.51 92.93a8 8 0 0 0-.68 6.03Z"/>`),
		g.Group(children),
	)
}