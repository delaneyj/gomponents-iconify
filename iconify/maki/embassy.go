package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Embassy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6.65 2C5.43 2 4.48 3.38 4.11 3.82a.49.49 0 0 0-.11.32v4.4a.44.44 0 0 0 .72.36a3 3 0 0 1 1.93-1.17C8.06 7.73 8.6 9 10.07 9a5.28 5.28 0 0 0 2.73-1.09a.49.49 0 0 0 .2-.4V2.45a.44.44 0 0 0-.62-.45a5.75 5.75 0 0 1-2.31 1.06C8.6 3.08 8.12 2 6.65 2zM2.5 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zM3 4v9.48a.5.5 0 0 1-1 0V4a.5.5 0 0 1 1 0z"/>`),
		g.Group(children),
	)
}