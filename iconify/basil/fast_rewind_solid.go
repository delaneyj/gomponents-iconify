package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastRewindSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.546 13.055a1.334 1.334 0 0 1 0-2.11A25.51 25.51 0 0 1 8.62 7.082l.466-.165c.87-.31 1.79.279 1.907 1.177c.138 1.05.22 2.119.244 3.19a1.34 1.34 0 0 1 .31-.339a25.51 25.51 0 0 1 7.073-3.863l.466-.165c.87-.31 1.79.279 1.907 1.177a30.31 30.31 0 0 1 0 7.812c-.118.898-1.037 1.486-1.907 1.178l-.466-.166a25.51 25.51 0 0 1-7.073-3.863a1.331 1.331 0 0 1-.31-.339a29.944 29.944 0 0 1-.244 3.19c-.118.898-1.037 1.486-1.907 1.178l-.466-.166a25.51 25.51 0 0 1-7.073-3.863Z"/>`),
		g.Group(children),
	)
}