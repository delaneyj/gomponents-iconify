package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NgButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M52 2H12C6.477 2 2 6.477 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.523-4.477-10-10-10zM30 33v9h-2.356l-14.31-14.313V42H10V22h2.356l14.311 14.309V22H30v11zm24-1.002C54 37.514 49.516 42 44 42c-5.513 0-10-4.486-10-10.002C34 26.484 38.487 22 44 22c2.566 0 5.007.969 6.87 2.732l-2.29 2.42a6.643 6.643 0 0 0-4.58-1.82a6.673 6.673 0 0 0-6.666 6.666A6.674 6.674 0 0 0 44 38.666a6.676 6.676 0 0 0 6.455-4.998H44v-3.336h10v1.666z"/>`),
		g.Group(children),
	)
}