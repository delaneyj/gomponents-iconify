package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignLeftRotateTwoHundredSeventyTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.25 18a.75.75 0 0 1-.75-.75V6.75a.75.75 0 0 1 1.5 0v10.5a.75.75 0 0 1-.75.75Zm5 0a.75.75 0 0 1-.75-.75V2.75a.75.75 0 0 1 1.5 0v14.5a.75.75 0 0 1-.75.75Zm4.25-.75a.75.75 0 0 0 1.5 0v-8.5a.75.75 0 0 0-1.5 0v8.5Z"/>`),
		g.Group(children),
	)
}