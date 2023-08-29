package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsPlayingSpadeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2H7c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2m0 18H7V4h10v16M12 7.7l-.6.5c-2 1.9-3.4 3.1-3.4 4.6C8 14 9 15 10.2 15c.4 0 .8-.1 1.2-.3l-.9 2.3h3l-.9-2.3c.3.2.8.3 1.2.3c1.2 0 2.2-.9 2.2-2.2c0-1.5-1.4-2.7-3.4-4.6l-.6-.5Z"/>`),
		g.Group(children),
	)
}