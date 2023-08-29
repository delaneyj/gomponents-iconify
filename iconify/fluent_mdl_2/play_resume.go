package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayResume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 256h128v1536H384V256zm1472 768L768 1792V256l1088 768zm-960 521l738-521l-738-521v1042z"/>`),
		g.Group(children),
	)
}