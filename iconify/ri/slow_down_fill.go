package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlowDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 13c0 2.21.895 4.21 2.343 5.657L4.93 20.07A9.969 9.969 0 0 1 2 13C2 7.477 6.477 3 12 3s10 4.477 10 10a9.97 9.97 0 0 1-2.929 7.071l-1.414-1.414A8 8 0 1 0 4 13Zm4.707-4.707L13.5 12.5l-2 2l-4.207-4.793l1.414-1.414Z"/>`),
		g.Group(children),
	)
}