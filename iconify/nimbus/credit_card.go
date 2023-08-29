package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.55 9.75h2V11h-2z"/><path fill="currentColor" d="M15.32 2.5H.68a.69.69 0 0 0-.68.71v9.58a.69.69 0 0 0 .68.71h14.64a.69.69 0 0 0 .68-.71V3.21a.69.69 0 0 0-.68-.71zm-.57 1.25V5H1.25V3.75zm-13.5 8.5V7.5h13.5v4.75z"/>`),
		g.Group(children),
	)
}