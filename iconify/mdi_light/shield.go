package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 3.11L19 6.63v5.01c0 4.81-3.22 9.26-7.5 10.4C7.22 20.9 4 16.45 4 11.64V6.63m7.5 16.44c4.9-1.23 8.5-6.13 8.5-11.43V6l-8.5-4L3 6v5.64c0 5.3 3.6 10.2 8.5 11.43Z"/>`),
		g.Group(children),
	)
}