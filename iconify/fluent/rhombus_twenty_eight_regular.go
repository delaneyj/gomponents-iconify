package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RhombusTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.555 5.928A3 3 0 0 1 9.357 4h14.638c2.103 0 3.553 2.107 2.802 4.072l-5.354 14A3 3 0 0 1 18.641 24H4.003C1.9 24 .45 21.893 1.2 19.928l5.354-14ZM9.357 5.5a1.5 1.5 0 0 0-1.401.964l-5.354 14a1.5 1.5 0 0 0 1.4 2.036h14.64a1.5 1.5 0 0 0 1.4-.964l5.354-14a1.5 1.5 0 0 0-1.4-2.036H9.355Z"/>`),
		g.Group(children),
	)
}