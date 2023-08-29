package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowheadDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaArrowheadDownOutline0"><g id="evaArrowheadDownOutline1"><g id="evaArrowheadDownOutline2" fill="currentColor"><path d="M17.37 12.39L12 16.71l-5.36-4.48a1 1 0 1 0-1.28 1.54l6 5a1 1 0 0 0 1.27 0l6-4.83a1 1 0 0 0 .15-1.41a1 1 0 0 0-1.41-.14Z"/><path d="M11.36 11.77a1 1 0 0 0 1.27 0l6-4.83a1 1 0 0 0 .15-1.41a1 1 0 0 0-1.41-.15L12 9.71L6.64 5.23a1 1 0 0 0-1.28 1.54Z"/></g></g></g>`),
		g.Group(children),
	)
}