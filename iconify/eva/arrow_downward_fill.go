package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownwardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaArrowDownwardFill0"><g id="evaArrowDownwardFill1"><path id="evaArrowDownwardFill2" fill="currentColor" d="M18.77 13.36a1 1 0 0 0-1.41-.13L13 16.86V5a1 1 0 0 0-2 0v11.86l-4.36-3.63a1 1 0 1 0-1.28 1.54l6 5l.15.09l.13.07a1 1 0 0 0 .72 0l.13-.07l.15-.09l6-5a1 1 0 0 0 .13-1.41Z"/></g></g>`),
		g.Group(children),
	)
}