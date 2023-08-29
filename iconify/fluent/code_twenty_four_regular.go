package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m8.066 18.943l6.5-14.5a.75.75 0 0 1 1.404.518l-.036.096l-6.5 14.5a.75.75 0 0 1-1.404-.518l.036-.096l6.5-14.5l-6.5 14.5ZM2.22 11.47l4.25-4.25a.75.75 0 0 1 1.133.976l-.073.085L3.81 12l3.72 3.719a.75.75 0 0 1-.976 1.133l-.084-.073l-4.25-4.25a.75.75 0 0 1-.073-.976l.073-.084l4.25-4.25l-4.25 4.25Zm14.25-4.25a.75.75 0 0 1 .976-.073l.084.073l4.25 4.25a.75.75 0 0 1 .073.976l-.073.085l-4.25 4.25a.75.75 0 0 1-1.133-.977l.073-.084L20.19 12l-3.72-3.72a.75.75 0 0 1 0-1.06Z"/>`),
		g.Group(children),
	)
}