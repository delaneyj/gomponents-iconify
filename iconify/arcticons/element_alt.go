package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElementAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.28 6a2.47 2.47 0 0 1 2.44-2.5a16.4 16.4 0 0 1 16.4 16.4h0a2.46 2.46 0 0 1-4.92 0A11.44 11.44 0 0 0 20.76 8.46A2.47 2.47 0 0 1 18.28 6Zm11.44 36a2.46 2.46 0 0 1-2.46 2.46a16.4 16.4 0 0 1-16.4-16.4a2.46 2.46 0 1 1 4.92 0h0a11.44 11.44 0 0 0 11.44 11.48a2.47 2.47 0 0 1 2.5 2.46ZM6 29.72a2.47 2.47 0 0 1-2.5-2.44a16.4 16.4 0 0 1 16.4-16.4a2.46 2.46 0 1 1 0 4.92h0A11.44 11.44 0 0 0 8.46 27.24A2.47 2.47 0 0 1 6 29.72Zm36-11.44a2.46 2.46 0 0 1 2.46 2.46a16.4 16.4 0 0 1-16.4 16.4a2.46 2.46 0 0 1 0-4.92a11.44 11.44 0 0 0 11.48-11.44a2.47 2.47 0 0 1 2.46-2.5Z"/>`),
		g.Group(children),
	)
}