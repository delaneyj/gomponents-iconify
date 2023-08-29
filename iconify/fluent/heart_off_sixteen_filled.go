package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartOffSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10.452 11.16l3.694 3.694a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708l1.946 1.945a3.25 3.25 0 0 0-.134 4.732l4.707 4.708a.5.5 0 0 0 .707 0l2.08-2.08Zm2.603-2.602l-1.188 1.188l-6.754-6.754a3.257 3.257 0 0 1 2.428.956l.454.453l.447-.448a3.252 3.252 0 0 1 4.6.012a3.25 3.25 0 0 1 .013 4.593Z"/>`),
		g.Group(children),
	)
}