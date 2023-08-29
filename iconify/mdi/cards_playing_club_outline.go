package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsPlayingClubOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2H7c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2m0 18H7V4h10v16M12 8c-1.1 0-2 .9-2 2c0 .4.1.7.3 1H10c-1.1 0-2 .9-2 2s.9 2 2 2c.6 0 1.1-.3 1.5-.7l-1 2.7h3l-1-2.7c.4.4.9.7 1.5.7c1.1 0 2-.9 2-2s-.9-2-2-2h-.3c.2-.3.3-.7.3-1c0-1.1-.9-2-2-2Z"/>`),
		g.Group(children),
	)
}