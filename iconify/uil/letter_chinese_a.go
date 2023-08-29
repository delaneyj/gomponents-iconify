package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterChineseA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 5h-7V4a1 1 0 0 0-2 0v1H4a1 1 0 0 0 0 2h11.882a14.493 14.493 0 0 1-3.94 7.952A14.426 14.426 0 0 1 8.663 9.67a1 1 0 0 0-1.889.66a16.414 16.414 0 0 0 3.68 5.958a14.299 14.299 0 0 1-5.768 2.735A1 1 0 0 0 4.899 21a1.018 1.018 0 0 0 .215-.023a16.297 16.297 0 0 0 6.831-3.319a16.387 16.387 0 0 0 6.842 3.319a1 1 0 0 0 .426-1.954a14.382 14.382 0 0 1-5.79-2.733A16.48 16.48 0 0 0 17.893 7H20a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}