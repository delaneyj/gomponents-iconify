package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 9a5 5 0 0 1 5-5h14a5 5 0 0 1 5 5v14a5 5 0 0 1-5 5H9a5 5 0 0 1-5-5V9Zm5-3a3 3 0 0 0-3 3v2h5V6H9Zm4 0v5h6V6h-6Zm0 7v6h6v-6h-6Zm-2 6v-6H6v6h5Zm-5 2v2a3 3 0 0 0 3 3h2v-5H6Zm7 0v5h6v-5h-6Zm8 0v5h2a3 3 0 0 0 3-3v-2h-5Zm5-2v-6h-5v6h5ZM21 6v5h5V9a3 3 0 0 0-3-3h-2Z"/>`),
		g.Group(children),
	)
}