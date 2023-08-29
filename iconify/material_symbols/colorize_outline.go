package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorizeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-4.75l8.95-8.95l-1.45-1.4l1.45-1.4l1.9 1.9l3.1-3.1q.275-.275.7-.275t.7.275l2.35 2.35q.275.275.275.7t-.275.7l-3.075 3.075l1.9 1.95L18.1 13.5l-1.4-1.45L7.75 21H3Zm2-2h1.95l8.3-8.35l-1.9-1.9L5 17.05V19ZM16.175 8.75l2.4-2.4l-.925-.925l-2.4 2.4l.925.925Zm0 0l-.925-.925l.925.925Z"/>`),
		g.Group(children),
	)
}