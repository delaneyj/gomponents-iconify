package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 9H9V7c0-.8.3-1.5.9-2.1c1.2-1.2 3.1-1.2 4.2 0c.4.4.6.9.8 1.4c.1.5.7.8 1.2.7c.5-.1.9-.7.7-1.2c-.2-.9-.7-1.7-1.3-2.3c-.9-1-2.2-1.5-3.5-1.5c-2.8 0-5 2.2-5 5v2c-1.7 0-3 1.3-3 3v7c0 1.7 1.3 3 3 3h10c1.7 0 3-1.3 3-3v-7c0-1.7-1.3-3-3-3zm-3.9 6.5l-.1.1V17c0 .6-.4 1-1 1s-1-.4-1-1v-1.4c-.6-.6-.7-1.5-.1-2.1c.6-.6 1.5-.7 2.1-.1c.6.5.7 1.5.1 2.1z"/>`),
		g.Group(children),
	)
}