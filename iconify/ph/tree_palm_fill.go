package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreePalmFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M239.84 60.33a8 8 0 0 1-4.65 5.75L179 90.55a71.42 71.42 0 0 1 43.36 33.21a70.64 70.64 0 0 1 7.2 54.32a8 8 0 0 1-12.56 4.28l-81-61.68V224a8 8 0 0 1-16 0V120.68l-81 61.68a8 8 0 0 1-12.57-4.28a70.64 70.64 0 0 1 7.2-54.32A71.42 71.42 0 0 1 77 90.55L20.81 66.08a8 8 0 0 1-2.6-12.85a66.86 66.86 0 0 1 97.74 0a72.21 72.21 0 0 1 12 17a72.21 72.21 0 0 1 12.05-17a66.86 66.86 0 0 1 97.74 0a8 8 0 0 1 2.1 7.1Z"/>`),
		g.Group(children),
	)
}