package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ankh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 12h-3.09A7.23 7.23 0 0 0 17 7A5 5 0 0 0 7 7a7.23 7.23 0 0 0 2.09 5H6a1 1 0 0 0 0 2h5v7a1 1 0 0 0 2 0v-7h5a1 1 0 0 0 0-2Zm-6-.16c-.93-.62-3-2.26-3-4.84a3 3 0 0 1 6 0c0 2.58-2.07 4.23-3 4.84Z"/>`),
		g.Group(children),
	)
}