package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTypeTicket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 0a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4Zm-.707 11.293l-.232-.232l-.829-.829L2 10l1.06-1.06l5.88-5.88L10 2l.232.232l.829.829l.232.232a1 1 0 0 0 1.415 1.414l.231.232l.829.829L14 6l-1.06 1.06l-5.88 5.88L6 14l-.232-.232l-.829-.829l-.232-.232a1 1 0 1 0-1.414-1.414Zm3.144.148A2.504 2.504 0 0 0 4.56 9.563l3-3.003l1.878 1.878l-3.002 3.002Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}