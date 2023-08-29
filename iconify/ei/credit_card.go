package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M39 38H11c-1.7 0-3-1.3-3-3V15c0-1.7 1.3-3 3-3h28c1.7 0 3 1.3 3 3v20c0 1.7-1.3 3-3 3zM11 14c-.6 0-1 .4-1 1v20c0 .6.4 1 1 1h28c.6 0 1-.4 1-1V15c0-.6-.4-1-1-1H11z"/><path fill="currentColor" d="M9 16h32v6H9zm3 10h1v2h-1zm2 0h1v2h-1zm2 0h1v2h-1zm3 0h1v2h-1zm2 0h1v2h-1zm2 0h1v2h-1zm3 0h1v2h-1zm2 0h1v2h-1zm2 0h1v2h-1zm3 0h1v2h-1zm2 0h1v2h-1zm2 0h1v2h-1z"/>`),
		g.Group(children),
	)
}