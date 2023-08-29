package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M453.122 79.012a128 128 0 0 0-181.087.068l-15.511 15.7l-15.382-15.666l-.1-.1a128 128 0 0 0-181.02 0l-6.91 6.91a128 128 0 0 0 0 181.019l182.373 182.371l20.595 21.578l.491-.492l.533.533l19.296-20.359L460.032 266.94a128.147 128.147 0 0 0 0-181.019ZM437.4 244.313L256.571 425.146L75.738 244.313a96 96 0 0 1 0-135.764l6.911-6.91a96 96 0 0 1 135.713-.051l38.093 38.787l38.274-38.736a96 96 0 0 1 135.765 0l6.91 6.909a96.11 96.11 0 0 1-.004 135.765Z"/>`),
		g.Group(children),
	)
}