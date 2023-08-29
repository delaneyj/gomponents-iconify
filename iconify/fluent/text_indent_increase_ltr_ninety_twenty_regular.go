package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentIncreaseLtrNinetyTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.854 2.646a.5.5 0 1 0-.708.708l1.5 1.5a.5.5 0 0 0 .708 0l1.5-1.5a.5.5 0 0 0-.708-.708L10 3.793L8.854 2.646ZM6 7.5a.5.5 0 0 0-1 0v7a.5.5 0 0 0 1 0v-7Zm9 0a.5.5 0 0 0-1 0v7a.5.5 0 0 0 1 0v-7Zm-5.5 10v-10a.5.5 0 0 1 1 0v10a.5.5 0 0 1-1 0Z"/>`),
		g.Group(children),
	)
}