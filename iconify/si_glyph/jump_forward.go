package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JumpForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.002 1c0-.553.442-1 .989-1H4.97c.547 0 .989.447.989 1v14c0 .553-.442 1-.989 1H2.991a.994.994 0 0 1-.989-1V1zm8.111 14.495c-.582.581-2.103.9-2.103-1.001V1.506c0-1.839 1.521-1.582 2.103-1l6.444 6.442a1.49 1.49 0 0 1 0 2.104l-6.444 6.443z"/>`),
		g.Group(children),
	)
}