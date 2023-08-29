package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SanitizerAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 8V5c0-.6-.4-1-1-1h-1V3h1c.6 0 1-.4 1-1s-.4-1-1-1h-4.8C8.5 1 6.9 2 6.1 3.6c-.2.4 0 1 .5 1.3c.5.2 1.1 0 1.3-.4c.4-.9 1.3-1.5 2.3-1.5H12v1h-1c-.6 0-1 .4-1 1v3c-1.7 0-3 1.3-3 3v9c0 1.7 1.3 3 3 3h6c1.7 0 3-1.3 3-3v-9c0-1.7-1.3-3-3-3zm-4-2h2v2h-2V6zm2 11h-2c-.6 0-1-.4-1-1s.4-1 1-1h2c.6 0 1 .4 1 1s-.4 1-1 1z"/>`),
		g.Group(children),
	)
}