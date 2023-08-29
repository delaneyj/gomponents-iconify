package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 85.46A16 16 0 0 0 224 80h-36.37L134 18.73a8 8 0 0 0-12 0L68.37 80H32a16 16 0 0 0-15.86 18.11L30 202.12A16 16 0 0 0 45.87 216h164.26A16 16 0 0 0 226 202.12l13.87-104A16 16 0 0 0 236 85.46ZM81.6 184a7.44 7.44 0 0 1-.81 0a8 8 0 0 1-8-7.2l-5.6-56a8 8 0 0 1 15.92-1.6l5.6 56a8 8 0 0 1-7.11 8.8Zm54.4-8a8 8 0 0 1-16 0v-56a8 8 0 0 1 16 0ZM89.63 80L128 36.15L166.37 80Zm99.13 40.8l-5.6 56a8 8 0 0 1-7.95 7.2a7.44 7.44 0 0 1-.81 0a8 8 0 0 1-7.16-8.76l5.6-56a8 8 0 0 1 15.92 1.6Z"/>`),
		g.Group(children),
	)
}