package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTrendingSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m8.75 4l4.548.001l.108.015l.113.033l.102.048l.038.023l.049.035l.076.067l.084.101l.059.1l.041.105l.023.097l.01.125v4.504a.75.75 0 0 1-1.493.102l-.006-.102L12.5 6.56l-4.22 4.223a.75.75 0 0 1-.978.073l-.084-.073L5.75 9.31l-2.469 2.47a.75.75 0 0 1-1.133-.977l.073-.084l2.998-3a.75.75 0 0 1 .977-.073l.084.073l1.47 1.474L11.441 5.5H8.75a.75.75 0 0 1-.743-.648L8 4.75A.75.75 0 0 1 8.75 4Z"/>`),
		g.Group(children),
	)
}