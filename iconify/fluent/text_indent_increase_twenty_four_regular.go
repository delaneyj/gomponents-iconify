package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentIncreaseTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.75 16a.75.75 0 0 1 0 1.5h-9a.75.75 0 0 1 0-1.5h9ZM2.72 9.22a.75.75 0 0 1 .976-.073l.084.073l2 2a.75.75 0 0 1 .073.976l-.073.084l-2 2a.75.75 0 0 1-1.133-.976l.073-.084l1.47-1.47l-1.47-1.47a.75.75 0 0 1 0-1.06ZM20.75 11a.75.75 0 0 1 0 1.5h-12a.75.75 0 0 1 0-1.5h12Zm-3-5a.75.75 0 0 1 0 1.5h-9a.75.75 0 0 1 0-1.5h9Z"/>`),
		g.Group(children),
	)
}