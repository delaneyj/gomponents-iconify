package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageRoundedX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 5.589 2 10c0 2.907 1.897 5.515 5 6.934V22l5.34-4.005C17.697 17.853 22 14.32 22 10c0-4.411-4.486-8-10-8zm3.707 10.293l-1.414 1.414L12 11.414l-2.293 2.293l-1.414-1.414L10.586 10L8.293 7.707l1.414-1.414L12 8.586l2.293-2.293l1.414 1.414L13.414 10l2.293 2.293z"/>`),
		g.Group(children),
	)
}