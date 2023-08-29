package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatStrikethroughVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.2 9.8c-1.2-2.3.5-5 2.9-5.5c3.1-1 7.6.4 7.5 4.2h-3c0-.3-.1-.6-.1-.8c-.2-.6-.6-.9-1.2-1.1c-.8-.3-2.1-.2-2.8.3C9 8.2 10.4 9.5 12 10H7.4c-.1-.1-.1-.2-.2-.2M21 13v-2H3v2h9.6c.2.1.4.1.6.2c.6.3 1.1.5 1.3 1.1c.1.4.2.9 0 1.3c-.2.5-.6.7-1.1.9c-1.8.5-4-.2-3.9-2.4h-3c-.1 2.6 2.1 4.4 4.5 4.7c3.8.8 8.3-1.6 6.3-5.9l3.7.1Z"/>`),
		g.Group(children),
	)
}