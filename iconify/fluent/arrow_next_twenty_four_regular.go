package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowNextTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.25 3a.75.75 0 0 1 .743.648L19 3.75v16.5a.75.75 0 0 1-1.493.102l-.007-.102V3.75a.75.75 0 0 1 .75-.75Zm-13.03.22a.75.75 0 0 1 .976-.073l.084.073l8.25 8.25a.75.75 0 0 1 .073.976l-.073.084l-8.25 8.25a.75.75 0 0 1-1.133-.976l.073-.084L12.94 12L5.22 4.28a.75.75 0 0 1 0-1.06Z"/>`),
		g.Group(children),
	)
}