package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckboxChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 19H7a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2ZM7 7v10h10V7H7Zm4 8.362l-2.7-2.647l1.4-1.43l1.3 1.271l3.3-3.267l1.4 1.422l-4.7 4.65v.001Z"/>`),
		g.Group(children),
	)
}