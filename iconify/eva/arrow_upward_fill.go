package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpwardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaArrowUpwardFill0"><g id="evaArrowUpwardFill1"><path id="evaArrowUpwardFill2" fill="currentColor" d="M5.23 10.64a1 1 0 0 0 1.41.13L11 7.14V19a1 1 0 0 0 2 0V7.14l4.36 3.63a1 1 0 1 0 1.28-1.54l-6-5l-.15-.09l-.13-.07a1 1 0 0 0-.72 0l-.13.07l-.15.09l-6 5a1 1 0 0 0-.13 1.41Z"/></g></g>`),
		g.Group(children),
	)
}