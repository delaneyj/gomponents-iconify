package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.1 5c2-.3 3.9-1.1 2.2-3.6c-.7-1-1.4-1.4-2-1.4c-1 0-1.7 1.1-2.3 2.2C7.4 1.1 6.7 0 5.7 0c-.6 0-1.3.4-2 1.4c-1.8 2.5.2 3.3 2.2 3.6H0v3h16V5h-5.9zm.2-4c.1 0 .5.1 1.2 1c.5.7.6 1.1.5 1.3c-.2.3-1.3.7-3.3.8c0-.2-.1-.4-.2-.6C9.1 2.1 9.8 1 10.3 1zM4 3.3c-.1-.2 0-.6.5-1.3c.7-.9 1.1-1 1.2-1c.5 0 1.2 1.1 1.8 2.5c-.1.2-.2.4-.2.6c-2-.1-3.1-.5-3.3-.8zM7 7V5h2v2H7zm2 8H7V9H1v7h14V9H9z"/>`),
		g.Group(children),
	)
}