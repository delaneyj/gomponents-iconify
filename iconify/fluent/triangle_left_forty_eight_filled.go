package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.228 27.528c-2.99-1.617-2.966-5.915.042-7.499L36.772 4.495c2.83-1.49 6.23.562 6.23 3.76v31.488c0 3.22-3.44 5.27-6.272 3.739L7.228 27.528Z"/>`),
		g.Group(children),
	)
}