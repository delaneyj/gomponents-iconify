package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFourFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Zm3.5-29.42V27.5h2.25a1.25 1.25 0 1 1 0 2.5H27.5v2.75a1.25 1.25 0 1 1-2.5 0V30h-8.442c-1.235 0-1.974-1.375-1.293-2.405l9.209-13.925c.902-1.364 3.026-.726 3.026.91ZM25 17.406V27.5h-6.675L25 17.406Z"/>`),
		g.Group(children),
	)
}