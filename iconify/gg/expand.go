package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.306 16.593l-.035 2l-6.999-.122l.123-7l2 .036l-.063 3.585l7.894-7.624l-3.532-.061l.035-2l6.999.122l-.123 7l-2-.036l.064-3.638l-7.948 7.676l3.585.062Z"/>`),
		g.Group(children),
	)
}