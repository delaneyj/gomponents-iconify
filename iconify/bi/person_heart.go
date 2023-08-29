package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-9 8c0 1 1 1 1 1h10s1 0 1-1s-1-4-6-4s-6 3-6 4Zm13.5-8.09c1.387-1.425 4.855 1.07 0 4.277c-4.854-3.207-1.387-5.702 0-4.276Z"/>`),
		g.Group(children),
	)
}