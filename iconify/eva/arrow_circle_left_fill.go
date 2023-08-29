package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaArrowCircleLeftFill0"><g id="evaArrowCircleLeftFill1"><path id="evaArrowCircleLeftFill2" fill="currentColor" d="M22 12a10 10 0 1 0-10 10a10 10 0 0 0 10-10Zm-11.86 3.69l-2.86-3a.49.49 0 0 1-.1-.15a.54.54 0 0 1-.1-.16a.94.94 0 0 1 0-.76a1 1 0 0 1 .21-.33l3-3a1 1 0 0 1 1.42 1.42L10.41 11H16a1 1 0 0 1 0 2h-5.66l1.25 1.31a1 1 0 0 1-1.45 1.38Z"/></g></g>`),
		g.Group(children),
	)
}