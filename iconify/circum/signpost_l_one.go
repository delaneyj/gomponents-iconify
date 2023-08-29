package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignpostLOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.437 16H6.536a2.489 2.489 0 0 1-1.744-.709L2.542 13.1a1.5 1.5 0 0 1 .007-2.2l2.243-2.191A2.483 2.483 0 0 1 6.536 8h13.9a1.5 1.5 0 0 1 1.5 1.5v5a1.5 1.5 0 0 1-1.499 1.5ZM6.536 9a1.491 1.491 0 0 0-1.046.425l-2.255 2.2a.5.5 0 0 0-.172.375a.494.494 0 0 0 .162.369l.01.01l2.254 2.2A1.492 1.492 0 0 0 6.536 15h13.9a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5Z"/>`),
		g.Group(children),
	)
}