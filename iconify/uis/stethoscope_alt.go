package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StethoscopeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.8 10c-.4-1.2-1.6-2-2.8-2c-1.7 0-3 1.3-3 3c0 1.3.8 2.4 2 2.8v1.7c0 2.5-2 4.5-4.5 4.5S9 18 9 15.5v-1l3.1-2.5c1.2-1 1.9-2.4 1.9-3.9V3c0-.6-.4-1-1-1h-2c-.6 0-1 .4-1 1s.4 1 1 1h1v4.1c0 .9-.4 1.8-1.1 2.3L8 12.7l-2.9-2.3C4.4 9.9 4 9 4 8.1V4h1c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1v5.1c0 1.5.7 3 1.9 3.9L7 14.5v1c0 3.6 2.9 6.5 6.5 6.5s6.5-2.9 6.5-6.5v-1.7c1.6-.5 2.4-2.2 1.8-3.8z"/>`),
		g.Group(children),
	)
}