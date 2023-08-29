package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PantoneLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.765 8l-.295-.73a1 1 0 0 1 .553-1.302l9.271-3.746a1 1 0 0 1 1.302.552l5.62 13.908a1 1 0 0 1-.553 1.302L12.39 21.73a1 1 0 0 1-1.302-.553L11 20.959v.04H7a1 1 0 0 1-1-1v-.269l-3.35-1.353a1 1 0 0 1-.553-1.302L5.765 8Zm2.236 11h2.208l-2.208-5.467V19Zm-2-6.244l-1.674 4.141l1.674.711v-4.852Zm1.698-5.309l4.87 12.054l7.417-2.997l-4.87-12.054L7.7 7.447Zm2.978 2.033a1 1 0 1 1-.75-1.855a1 1 0 0 1 .75 1.855Z"/>`),
		g.Group(children),
	)
}