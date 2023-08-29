package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOffSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12.274 12.981l1.872 1.873a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708L2.74 3.447C2.288 3.814 2 4.373 2 5v6a2 2 0 0 0 2 2h8a1.9 1.9 0 0 0 .274-.019ZM11.293 12H4a1 1 0 0 1-1-1V6.876L7.763 9.44a.5.5 0 0 0 .474 0l.323-.173L11.293 12ZM7.404 8.111L3 5.74V5a1 1 0 0 1 .455-.838l3.95 3.95ZM13 5.74L9.66 7.538l.738.739L13 6.876v4.003l.86.86c.09-.23.14-.478.14-.739V5a2 2 0 0 0-2-2H5.121l1 1H12a1 1 0 0 1 1 1v.74Z"/>`),
		g.Group(children),
	)
}