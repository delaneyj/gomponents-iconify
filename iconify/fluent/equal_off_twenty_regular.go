package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualOffTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.854 2.146a.5.5 0 1 0-.708.708L6.293 7H3.5a.5.5 0 0 0 0 1h3.793l4 4H3.5a.5.5 0 0 0 0 1h8.793l4.853 4.854a.5.5 0 0 0 .708-.708l-15-15ZM14.12 12l1 1h1.38a.5.5 0 0 0 0-1h-2.379Zm-5-5l1 1h6.38a.5.5 0 0 0 0-1H9.121Z"/>`),
		g.Group(children),
	)
}