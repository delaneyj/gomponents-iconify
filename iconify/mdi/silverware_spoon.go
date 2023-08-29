package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SilverwareSpoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.88 11.53l-9.76 9.76l-1.41-1.41l9.76-9.76c-.71-1.53-.21-3.68 1.38-5.27c1.91-1.92 4.65-2.28 6.11-.82c1.47 1.47 1.11 4.21-.81 6.12c-1.59 1.59-3.74 2.09-5.27 1.38Z"/>`),
		g.Group(children),
	)
}