package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagVariantOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.8 22.7L10.2 12.1C7.8 12.4 7 14 7 14v7H5V6.9L1.1 3l1.3-1.3l19.7 19.7l-1.3 1.3M20 12V4s-1 2-4 2c-2 0-2-2-5-2c-1.2 0-2.3.3-3.2.6l9.3 9.3C19.2 13.5 20 12 20 12Z"/>`),
		g.Group(children),
	)
}