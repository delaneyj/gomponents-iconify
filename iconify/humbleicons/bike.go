package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.5 10.5l-.6-.8a1 1 0 0 0 .153 1.694L8.5 10.5Zm4 2l.99.141a1 1 0 0 0-.543-1.035l-.447.894Zm-1.49 3.359a1 1 0 1 0 1.98.282l-1.98-.282ZM12.5 7.5l.707-.707A1 1 0 0 0 11.9 6.7l.6.8ZM15 10l-.707.707A1 1 0 0 0 15 11v-1Zm2 1a1 1 0 1 0 0-2v2Zm-9 5a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4H8Zm-2 2a2 2 0 0 1-2-2H2a4 4 0 0 0 4 4v-2Zm-2-2a2 2 0 0 1 2-2v-2a4 4 0 0 0-4 4h2Zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4v2Zm2.053-2.606l4 2l.894-1.788l-4-2l-.894 1.788Zm3.457.965l-.5 3.5l1.98.282l.5-3.5l-1.98-.282ZM9.1 11.3l4-3l-1.2-1.6l-4 3l1.2 1.6Zm2.693-3.093l2.5 2.5l1.414-1.414l-2.5-2.5l-1.414 1.414ZM15 11h2V9h-2v2Zm5 5a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4h-2Zm-2 2a2 2 0 0 1-2-2h-2a4 4 0 0 0 4 4v-2Zm-2-2a2 2 0 0 1 2-2v-2a4 4 0 0 0-4 4h2Zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4v2Zm-3-9v2a2 2 0 0 0 2-2h-2Zm0 0h-2a2 2 0 0 0 2 2V5Zm0 0V3a2 2 0 0 0-2 2h2Zm0 0h2a2 2 0 0 0-2-2v2Z"/>`),
		g.Group(children),
	)
}