package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zincsearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.723 19.748l-1.73 1.493H.678L0 18.77l10.63-9.343l.542 6.635h8.701a3.649 3.649 0 0 1-1.15 3.686zM5.277 4.252l1.73-1.493h16.316L24 5.23l-10.63 9.343l-.542-6.635H4.129a3.648 3.648 0 0 1 1.148-3.686Z"/>`),
		g.Group(children),
	)
}