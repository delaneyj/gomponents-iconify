package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.09 4.5H19a2.05 2.05 0 0 1 2.05 2.05V17.2c0 1.72-1.91 2.74-3 1.65l-3.48-3.5A12.82 12.82 0 1 0 36.83 24c0-3.46-1-6.06-4.93-10c-.84-.84-.1-1.55.38-2l2.28-2.27c.94-.94 1.27-1.12 2.1-.29c4.86 4.86 6.83 8.52 6.84 14.63A19.5 19.5 0 1 1 9.86 10.63C8.69 9.44 8.32 9.06 6.35 7.09c-.77-.76-.08-2.59 1.74-2.59Z"/>`),
		g.Group(children),
	)
}