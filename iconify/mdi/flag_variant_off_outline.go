package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagVariantOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.1 21.5L2.4 1.7L1.1 3L5 6.9V21h2v-7s.8-1.6 3.2-1.9l10.7 10.7l1.2-1.3M7 11.5V8.9l1.7 1.7c-1 .4-1.7.9-1.7.9m2.4-5.3L7.8 4.6c.9-.3 2-.6 3.2-.6c3 0 3 2 5 2c3 0 4-2 4-2v8s-.8 1.5-2.9 1.9L15 11.8c.3.1.6.2 1 .2c2 0 2-1 2-1V7.5S17 8 16 8c-2 0-3-2-5-2c-.5 0-1.1.1-1.6.2Z"/>`),
		g.Group(children),
	)
}