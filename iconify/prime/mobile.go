package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 3.25H8c-.96 0-1.75.79-1.75 1.75v14c0 .96.79 1.75 1.75 1.75h8c.96 0 1.75-.79 1.75-1.75V5c0-.96-.79-1.75-1.75-1.75ZM16.25 19c0 .14-.11.25-.25.25H8c-.14 0-.25-.11-.25-.25V5c0-.14.11-.25.25-.25h8c.14 0 .25.11.25.25v14ZM13 16c0 .55-.45 1-1 1s-1-.45-1-1s.45-1 1-1s1 .45 1 1Z"/>`),
		g.Group(children),
	)
}