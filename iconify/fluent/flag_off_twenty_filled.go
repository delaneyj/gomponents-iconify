package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOffTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.854 2.146a.5.5 0 1 0-.708.708L3.5 4.207V17.5a.5.5 0 0 0 1 0V13h7.793l4.853 4.854a.5.5 0 0 0 .708-.708L3.765 3.058l-.911-.912ZM15.5 13h-.379l-10-10H15.5a.5.5 0 0 1 .416.777L13.101 8l2.815 4.223A.5.5 0 0 1 15.5 13Z"/>`),
		g.Group(children),
	)
}