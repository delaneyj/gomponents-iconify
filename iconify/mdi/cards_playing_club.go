package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsPlayingClub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2H7c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2m-3 13c-.6 0-1.1-.3-1.5-.7l1 2.7h-3l1-2.7c-.4.4-.9.7-1.5.7c-1.1 0-2-.9-2-2s.9-2 2-2h.3c-.2-.3-.3-.7-.3-1c0-1.1.9-2 2-2s2 .9 2 2c0 .4-.1.7-.3 1h.3c1.1 0 2 .9 2 2s-.9 2-2 2Z"/>`),
		g.Group(children),
	)
}