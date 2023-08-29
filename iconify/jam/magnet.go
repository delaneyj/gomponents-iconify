package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 12c0-2.878-1.516-5-3-5s-3 2.122-3 5v6a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-6C0 5.373 4.477 0 10 0s10 5.373 10 12v6a2 2 0 0 1-2 2h-3a2 2 0 0 1-2-2v-6zm2 0h3c0-5.595-3.67-10-8-10S2 6.405 2 12h3c0-3.866 2.239-7 5-7s5 3.134 5 7zM2 18h3v-4H2v4zm13 0h3v-4h-3v4z"/>`),
		g.Group(children),
	)
}