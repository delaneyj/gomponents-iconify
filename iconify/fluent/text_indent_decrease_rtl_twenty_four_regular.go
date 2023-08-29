package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentDecreaseRtlTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 16.75a.75.75 0 0 1 .75-.75h9a.75.75 0 0 1 0 1.5h-9a.75.75 0 0 1-.75-.75Zm-3-5a.75.75 0 0 1 .75-.75h12a.75.75 0 0 1 0 1.5h-12a.75.75 0 0 1-.75-.75Zm3-5A.75.75 0 0 1 6.25 6h9a.75.75 0 0 1 0 1.5h-9a.75.75 0 0 1-.75-.75Zm13.696 2.397a.75.75 0 0 0-.976 1.133l1.47 1.47l-1.47 1.47l-.073.084a.75.75 0 0 0 1.133.976l2-2l.073-.084a.75.75 0 0 0-.073-.976l-2-2l-.084-.073Z"/>`),
		g.Group(children),
	)
}