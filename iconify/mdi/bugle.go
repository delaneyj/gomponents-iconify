package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bugle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 6c0 5-9 5-9 5H5c-1 0-2-1-2-1H2v4h1s1-1 2-1h1.3c-.8.5-1.3 1.2-1.3 2c0 1.8 2.3 3 5.5 3s5.5-1.2 5.5-3c0-.6-.3-1.2-.8-1.7c2.6.5 5.8 1.7 5.8 4.7h1V6h-1M10.5 16.7c-2.3 0-4.1-.8-4.1-1.7c0-.9 1.8-1.7 4.1-1.7s4.1.8 4.1 1.7c0 .9-1.8 1.7-4.1 1.7Z"/>`),
		g.Group(children),
	)
}