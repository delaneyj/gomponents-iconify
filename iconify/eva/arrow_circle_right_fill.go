package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaArrowCircleRightFill0"><g id="evaArrowCircleRightFill1"><path id="evaArrowCircleRightFill2" fill="currentColor" d="M2 12A10 10 0 1 0 12 2A10 10 0 0 0 2 12Zm11.86-3.69l2.86 3a.49.49 0 0 1 .1.15a.54.54 0 0 1 .1.16a.94.94 0 0 1 0 .76a1 1 0 0 1-.21.33l-3 3a1 1 0 0 1-1.42-1.42l1.3-1.29H8a1 1 0 0 1 0-2h5.66l-1.25-1.31a1 1 0 0 1 1.45-1.38Z"/></g></g>`),
		g.Group(children),
	)
}