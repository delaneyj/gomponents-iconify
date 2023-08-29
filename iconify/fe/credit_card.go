package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCreditCard0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCreditCard1" fill="currentColor" fill-rule="nonzero"><path id="feCreditCard2" d="M4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm0 2v3h16V6H4Zm0 5v7h16v-7H4Zm2 3v2h4v-2H6Zm6 0v2h4v-2h-4Z"/></g></g>`),
		g.Group(children),
	)
}