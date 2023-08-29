package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDownFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20S35.046 4 24 4Zm-8.634 19.866a1.25 1.25 0 0 1 1.768 0l5.616 5.616V15.25a1.25 1.25 0 1 1 2.5 0v14.232l5.616-5.616a1.25 1.25 0 0 1 1.768 1.768l-7.75 7.75a1.25 1.25 0 0 1-1.768 0l-7.75-7.75a1.25 1.25 0 0 1 0-1.768Z"/>`),
		g.Group(children),
	)
}