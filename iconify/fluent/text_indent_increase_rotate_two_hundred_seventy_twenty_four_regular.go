package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentIncreaseRotateTwoHundredSeventyTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 6.25a.75.75 0 0 1 1.5 0v9a.75.75 0 0 1-1.5 0v-9ZM9.22 21.28a.75.75 0 0 1-.073-.976l.073-.084l2-2a.75.75 0 0 1 .976-.073l.084.073l2 2a.75.75 0 0 1-.976 1.133l-.084-.073l-1.47-1.47l-1.47 1.47a.75.75 0 0 1-1.06 0ZM11 3.25a.75.75 0 0 1 1.5 0v12a.75.75 0 0 1-1.5 0v-12Zm-5 3a.75.75 0 0 1 1.5 0v9a.75.75 0 0 1-1.5 0v-9Z"/>`),
		g.Group(children),
	)
}