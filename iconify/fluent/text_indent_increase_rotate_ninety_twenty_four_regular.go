package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentIncreaseRotateNinetyTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 17.75a.75.75 0 0 1-1.5 0v-9a.75.75 0 0 1 1.5 0v9Zm6.78-15.03a.75.75 0 0 1 .073.976l-.073.084l-2 2a.75.75 0 0 1-.976.073l-.084-.073l-2-2a.75.75 0 0 1 .976-1.133l.084.073l1.47 1.47l1.47-1.47a.75.75 0 0 1 1.06 0ZM13 20.75a.75.75 0 0 1-1.5 0v-12a.75.75 0 0 1 1.5 0v12Zm5-3a.75.75 0 0 1-1.5 0v-9a.75.75 0 0 1 1.5 0v9Z"/>`),
		g.Group(children),
	)
}