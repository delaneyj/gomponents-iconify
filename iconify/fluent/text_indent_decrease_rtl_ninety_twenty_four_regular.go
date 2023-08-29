package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentDecreaseRtlNinetyTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.25 5.5a.75.75 0 0 1 .75.75v9a.75.75 0 0 1-1.5 0v-9a.75.75 0 0 1 .75-.75Zm5-3a.75.75 0 0 1 .75.75v12a.75.75 0 0 1-1.5 0v-12a.75.75 0 0 1 .75-.75Zm5 3a.75.75 0 0 1 .75.75v9a.75.75 0 0 1-1.5 0v-9a.75.75 0 0 1 .75-.75Zm-2.397 13.696a.75.75 0 0 0-1.133-.976l-1.47 1.47l-1.47-1.47l-.084-.073a.75.75 0 0 0-.976 1.133l2 2l.084.073a.75.75 0 0 0 .976-.073l2-2l.073-.084Z"/>`),
		g.Group(children),
	)
}