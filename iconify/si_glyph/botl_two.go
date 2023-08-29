package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BotlTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 0h1.969v2.03H8zm1.886 3H8.06s.497 1.247-1.112 2.006C5.691 5.599 6 5.259 6 7.203V14.9c0 .608.521 1.1 1.16 1.1h3.679c.64 0 1.161-.492 1.161-1.1V7.203c-.001-2.029.187-1.436-.97-2.166C9.505 4.08 9.886 3 9.886 3zM11 15h-1V7h1v8z"/>`),
		g.Group(children),
	)
}