package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LibraryTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.99 3a2 2 0 0 1 2 2v18a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h1.99Zm7 0a2 2 0 0 1 2 2v18a2 2 0 0 1-2 2h-1.995a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h1.995Zm9.08 3.543l3.86 15.483a2.002 2.002 0 0 1-1.457 2.425l-1.963.49a2 2 0 0 1-2.424-1.458L16.226 8a2 2 0 0 1 1.456-2.425l1.963-.489a2 2 0 0 1 2.425 1.457Z"/>`),
		g.Group(children),
	)
}