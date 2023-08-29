package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicRestSixteenth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 3a5.56 5.56 0 0 1-3.05 1.86A1.5 1.5 0 1 0 10.5 6h.24a6.32 6.32 0 0 0 3.51-1.07L12.9 9.1a5.56 5.56 0 0 1-2.95 1.76A1.5 1.5 0 1 0 8.5 12h.24a6.32 6.32 0 0 0 3.51-1.07L9 21h2l6-18Z"/>`),
		g.Group(children),
	)
}