package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.927 23.935a2.4 2.4 0 0 1-1.882-1.883a2.4 2.4 0 0 1 2.82-2.82a2.4 2.4 0 0 1 1.882 1.882a2.4 2.4 0 0 1-2.82 2.82zM21.135 4.768a2.4 2.4 0 0 1-1.882-1.882a2.4 2.4 0 0 1 2.82-2.82a2.4 2.4 0 0 1 1.882 1.882a2.4 2.4 0 0 1-2.82 2.82zM12.021 6.01c1.147 0 2.219.324 3.13.883l2.585-2.585A9.583 9.583 0 0 0 4.328 17.716l2.586-2.586a5.99 5.99 0 0 1 5.107-9.12zm5.107 2.86a5.99 5.99 0 0 1-8.237 8.237l-2.586 2.585A9.583 9.583 0 0 0 19.713 6.284Z"/>`),
		g.Group(children),
	)
}