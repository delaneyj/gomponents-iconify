package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Power(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13V4h1v9h-1m8-.5a7.5 7.5 0 0 1-7.5 7.5A7.5 7.5 0 0 1 4 12.5c0-2.71 1.44-5.09 3.6-6.4l.73.73A6.478 6.478 0 0 0 5 12.5a6.5 6.5 0 0 0 6.5 6.5a6.5 6.5 0 0 0 6.5-6.5c0-2.44-1.34-4.56-3.33-5.67l.73-.73a7.476 7.476 0 0 1 3.6 6.4Z"/>`),
		g.Group(children),
	)
}