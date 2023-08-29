package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.528 2.973a.75.75 0 0 1 .472.696v8.662a.75.75 0 0 1-.472.696l-7.25 2.9a.75.75 0 0 1-.557 0l-7.25-2.9A.75.75 0 0 1 0 12.331V3.669a.75.75 0 0 1 .471-.696L7.443.184l.004-.001l.274-.11a.75.75 0 0 1 .558 0l.274.11l.004.001l6.971 2.789Zm-1.374.527L8 5.962L1.846 3.5L1 3.839v.4l6.5 2.6v7.922l.5.2l.5-.2V6.84l6.5-2.6v-.4l-.846-.339Z"/>`),
		g.Group(children),
	)
}