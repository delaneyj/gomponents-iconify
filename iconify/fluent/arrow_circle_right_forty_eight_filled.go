package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleRightFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 24C4 12.954 12.954 4 24 4s20 8.954 20 20s-8.954 20-20 20S4 35.046 4 24Zm19.866-8.634a1.25 1.25 0 0 0 0 1.768l5.616 5.616H15.25a1.25 1.25 0 1 0 0 2.5h14.232l-5.616 5.616a1.25 1.25 0 0 0 1.768 1.768l7.75-7.75a1.25 1.25 0 0 0 0-1.768l-7.75-7.75a1.25 1.25 0 0 0-1.768 0Z"/>`),
		g.Group(children),
	)
}