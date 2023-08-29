package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageLockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 3v-.5C23 1.1 21.9 0 20.5 0S18 1.1 18 2.5V3c-.5 0-1 .5-1 1v4c0 .5.5 1 1 1h5c.5 0 1-.5 1-1V4c0-.5-.5-1-1-1m-1 0h-3v-.5c0-.8.7-1.5 1.5-1.5s1.5.7 1.5 1.5V3m0 8v5c0 1.1-.9 2-2 2H6l-4 4V4c0-1.1.9-2 2-2h11v2H4v13.2L5.2 16H20v-5h2Z"/>`),
		g.Group(children),
	)
}