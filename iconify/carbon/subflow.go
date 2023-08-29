package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 30A14.016 14.016 0 0 1 2 16h2a12.002 12.002 0 0 0 18.46 10.114l1.08 1.684A13.951 13.951 0 0 1 16 30Z"/><path fill="currentColor" d="m26 17l-1.414 1.414L26.172 20H15.816a2.987 2.987 0 0 0-.275-.576l4.481-5.601A2.968 2.968 0 0 0 21 14a3 3 0 1 0-2.816-4h-4.368a2.982 2.982 0 0 0-5.632 0H2v2h6.184a2.982 2.982 0 0 0 5.632 0h4.368a2.987 2.987 0 0 0 .274.576l-4.48 5.601A2.968 2.968 0 0 0 13 18a3 3 0 1 0 2.816 4h10.356l-1.586 1.586L26 25l4-4Zm-5-7a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm-10 2a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm2 10a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><path fill="currentColor" d="M30 16h-2A12.002 12.002 0 0 0 9.54 5.886L8.46 4.202A14.002 14.002 0 0 1 30 16Z"/>`),
		g.Group(children),
	)
}