package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListOrder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feListOrder0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feListOrder1" fill="currentColor"><path id="feListOrder2" d="M10 4h10a1 1 0 0 1 0 2H10a1 1 0 1 1 0-2Zm0 7h10a1 1 0 0 1 0 2H10a1 1 0 0 1 0-2Zm0 7h10a1 1 0 0 1 0 2H10a1 1 0 0 1 0-2ZM4.5 4a.5.5 0 0 1 0-1h1c.28 0 .5.22.5.5v3a.5.5 0 0 1-1 0V4.015L4.5 4Zm0 6h2c.28 0 .5.22.5.5v.5l-1.7 2h1.2c.28 0 .5.22.5.5s-.224.5-.5.5h-2c-.28 0-.5-.22-.5-.5V13l1.7-2H4.5c-.28 0-.5-.22-.5-.5s.199-.5.5-.5Zm2 11h-2a.5.5 0 1 1 0-1H6v-.5H4.5a.5.5 0 1 1 0-1H6V18H4.5a.5.5 0 1 1 0-1h2a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5Z"/></g></g>`),
		g.Group(children),
	)
}