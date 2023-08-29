package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlusSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5 3.4c-.8 0-1.3.8-1.2 1.8c.1 1.1.9 1.9 1.7 2c.8 0 1.3-.8 1.2-1.9c-.1-1-.9-1.9-1.7-1.9zm.4 5.9c-1.2 0-2.3.7-2.3 1.6s.9 1.7 2.1 1.7c1.7 0 2.3-.7 2.3-1.6v-.3c-.1-.5-.6-.8-1.3-1.2c-.2-.2-.5-.2-.8-.2z"/><path fill="currentColor" d="M0 0v16h16V0H0zm7.9 5.3c0 .7-.4 1.2-.9 1.6s-.6.6-.6.9c0 .3.5.8.8 1c.8.6 1.1 1.1 1.1 2c0 1.1-1.1 2.3-3.1 2.3c-1.7 0-3.2-.7-3.2-1.8C2 10.1 3.3 9 5.1 9h.5c-.2-.3-.4-.6-.4-.9c0-.2.1-.4.2-.6h-.3c-1.4 0-2.4-1-2.4-2.3C2.7 4 4 2.9 5.4 2.9h3.1l-.7.6h-1c.7.2 1.1 1 1.1 1.8zm6.1.2h-2.1v2h-.5v-2h-2V5h2V3h.5v2H14v.5z"/>`),
		g.Group(children),
	)
}