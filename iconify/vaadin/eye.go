package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 3.9C1.3 3.9 0 9 0 9s2.2 4.1 7.9 4.1s8.1-4 8.1-4s-1.3-5.2-8-5.2zM5.3 5.4c.5-.3 1.3-.3 1.3-.3s-.5.9-.5 1.6c0 .7.2 1.1.2 1.1L5.2 8s-.3-.5-.3-1.2c0-.8.4-1.4.4-1.4zm2.6 6.7c-4.1 0-6.2-2.3-6.8-3.2c.3-.7 1.1-2.2 3.1-3.2c-.1.4-.2.8-.2 1.3c0 2.2 1.8 4 4 4s4-1.8 4-4c0-.5-.1-.9-.2-1.3c2 .9 2.8 2.5 3.1 3.2c-.7.9-2.8 3.2-7 3.2z"/>`),
		g.Group(children),
	)
}