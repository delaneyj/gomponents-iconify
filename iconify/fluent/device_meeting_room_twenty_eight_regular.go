package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceMeetingRoomTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.35 3a3.5 3.5 0 0 0-3.378 2.58L2.143 15.946C1.45 18.492 3.365 21 6.003 21h15.994c2.638 0 4.554-2.508 3.86-5.053L23.028 5.58A3.5 3.5 0 0 0 19.651 3H8.349ZM6.42 5.974A2 2 0 0 1 8.35 4.5h11.3a2 2 0 0 1 1.93 1.474l2.83 10.368a2.5 2.5 0 0 1-2.413 3.158H6.003a2.5 2.5 0 0 1-2.412-3.158L6.419 5.974ZM7.747 22.5a.75.75 0 1 0 0 1.5H20.25a.75.75 0 1 0 0-1.5H7.748Z"/>`),
		g.Group(children),
	)
}