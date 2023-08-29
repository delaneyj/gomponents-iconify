package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAutofitUpTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m15.712 5.233l-.71-.735V17.25a.75.75 0 1 1-1.5 0V4.494l-.713.739A.75.75 0 0 1 11.71 4.19l1.821-1.886a1 1 0 0 1 1.44 0l1.82 1.886a.75.75 0 0 1-1.079 1.042ZM5 3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h6.25a.75.75 0 0 0 0-1.5H5a.5.5 0 0 1-.5-.5V5a.5.5 0 0 1 .5-.5h4.25a.75.75 0 0 0 0-1.5H5Z"/>`),
		g.Group(children),
	)
}