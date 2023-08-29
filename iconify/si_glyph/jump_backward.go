package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JumpBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.994 1c0-.553-.442-1-.989-1h-1.979a.994.994 0 0 0-.99 1v14c0 .553.443 1 .99 1h1.979a.994.994 0 0 0 .989-1V1zM.438 9.052a1.49 1.49 0 0 1 0-2.104L6.882.506c.581-.582 2.103-.839 2.103 1v12.988c0 1.901-1.521 1.582-2.103 1.001L.438 9.052z"/>`),
		g.Group(children),
	)
}