package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentDecreaseRtlSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 3.75A.75.75 0 0 1 4.75 3h5.5a.75.75 0 0 1 0 1.5h-5.5A.75.75 0 0 1 4 3.75Zm8.22 1.97a.75.75 0 0 1 1.06 0l1.5 1.5a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 1 1-1.06-1.06l.97-.97l-.97-.97a.75.75 0 0 1 0-1.06ZM1 7.75A.75.75 0 0 1 1.75 7h8.5a.75.75 0 0 1 0 1.5h-8.5A.75.75 0 0 1 1 7.75Zm5 4a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}