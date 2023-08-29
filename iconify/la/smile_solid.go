package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4C9.383 4 4 9.383 4 16s5.383 12 12 12s12-5.383 12-12S22.617 4 16 4zm0 2c5.535 0 10 4.465 10 10s-4.465 10-10 10S6 21.535 6 16S10.465 6 16 6zm-4.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm9 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm-9.688 7l-1.718 1c1.383 2.387 3.953 4 6.906 4s5.523-1.613 6.906-4l-1.718-1A5.977 5.977 0 0 1 16 22a5.977 5.977 0 0 1-5.188-3z"/>`),
		g.Group(children),
	)
}