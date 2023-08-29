package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RatingMatureSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.5A2.5 2.5 0 0 1 4.5 2h7A2.5 2.5 0 0 1 14 4.5v7a2.5 2.5 0 0 1-2.5 2.5h-7A2.5 2.5 0 0 1 2 11.5v-7Zm3.924.735A.5.5 0 0 0 5 5.5v5a.5.5 0 1 0 1 0V7.243l1.576 2.522a.5.5 0 0 0 .848 0L10 7.243V10.5a.5.5 0 1 0 1 0v-5a.5.5 0 0 0-.924-.265L8 8.557L5.924 5.235Z"/>`),
		g.Group(children),
	)
}