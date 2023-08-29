package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowClockwiseTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 2A.75.75 0 0 0 9 2v.646a4.5 4.5 0 1 0 1.42 4.206c.088-.465-.304-.852-.777-.852c-.355 0-.636.291-.711.638a3.001 3.001 0 1 1-1.266-3.133H7.25a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 .75-.75V2Z"/>`),
		g.Group(children),
	)
}