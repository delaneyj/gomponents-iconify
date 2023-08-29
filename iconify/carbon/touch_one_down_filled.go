package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TouchOneDownFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 30a6.007 6.007 0 0 1-6-6h2a4 4 0 0 0 8 0h2a6.007 6.007 0 0 1-6 6Z"/><path fill="currentColor" d="M21 2h-4.44a4 4 0 0 0-2.708 1.057l-9.2 8.466a2.002 2.002 0 0 0 .118 3.055a2.074 2.074 0 0 0 2.658-.173L11 11.143V24a2 2 0 0 0 4 0v-7a2 2 0 0 0 4 0v-1a2 2 0 0 0 4 0v-1a2 2 0 0 0 4 0V8a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}