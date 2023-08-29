package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageQuestion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v18l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2m-7 12h-2v-2h2v2m1.8-5c-.3.4-.7.6-1.1.8c-.3.2-.4.3-.5.5c-.2.2-.2.4-.2.7h-2c0-.5.1-.8.3-1.1c.2-.2.6-.5 1.1-.8c.3-.1.5-.3.6-.5c.1-.2.2-.5.2-.7c0-.3-.1-.5-.3-.7c-.2-.2-.5-.3-.8-.3c-.3 0-.5.1-.7.2c-.2.1-.3.3-.3.6h-2c.1-.7.4-1.3.9-1.7c.5-.4 1.2-.5 2.1-.5c.9 0 1.7.2 2.2.6c.5.4.8 1 .8 1.7c.1.4 0 .8-.3 1.2Z"/>`),
		g.Group(children),
	)
}