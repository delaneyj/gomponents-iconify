package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOffSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m9 9.707l5.146 5.147a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708L6.293 7L3.338 9.955a1.65 1.65 0 0 0-.398.644l-.914 2.743a.5.5 0 0 0 .632.632l2.743-.914c.243-.08.463-.217.644-.398L9 9.707Zm3-2.999l-1.586 1.585l-2.707-2.707l1.585-1.585L12 6.708ZM10.732 2.56a1.914 1.914 0 0 1 2.707 2.708L12.707 6L9.998 3.294l.734-.734Z"/>`),
		g.Group(children),
	)
}