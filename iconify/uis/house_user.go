package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.7 10.3l-9-8c-.4-.3-.9-.3-1.3 0l-9 8c-.4.4-.5 1-.1 1.4s1 .5 1.4.1l.3-.4V21c0 .6.4 1 1 1h14c.6 0 1-.4 1-1v-9.6l.3.3c.4.4 1 .3 1.4-.1c.4-.3.4-1 0-1.3zM12 11c1.5 0 2.7 1.2 2.7 2.7c0 1.5-1.2 2.7-2.7 2.7c-1.5 0-2.7-1.2-2.7-2.7S10.5 11 12 11zm-5 9c0-.1 0-.1.1-.2c2.2-2.7 6.2-3.2 8.9-1c.4.3.7.6 1 1c0 0 0 .1.1.2H7z"/>`),
		g.Group(children),
	)
}