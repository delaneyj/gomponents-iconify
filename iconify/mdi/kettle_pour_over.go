package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KettlePourOver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 4c-.6 0-1 .4-1 1v1h5V5c0-.6-.4-1-1-1h-3M3 7l1.1 7.1c.1.7.4 1.4 1 2c.4.5 1.1.8 1.9.9c0 .6.4 1 1 1h8c.6 0 1-.4 1-1l-2-6h2l3.3 5l1.7-1l-3.3-5l1.3-.8l-1-1.7L16.4 9H15V7H8v4l-.8 5c-.7-.1-1.1-.3-1.4-.6c-.4-.4-.6-.9-.7-1.5L4 7H3m1 12v2h16v-2H4Z"/>`),
		g.Group(children),
	)
}