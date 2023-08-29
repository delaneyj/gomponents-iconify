package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LighthouseOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 10V2l9 3v2l-9 3m7 0V8h5V4H8V3l4-2l4 2v1h-1v4h1v2h-1.26l-6.3 3.64L9 10H8M7 23l.04-.24l9.11-5.26l.52 3.38L13 23H7m1.05-6.83L15.31 12l.52 3.37l-8.4 4.85l.62-4.05Z"/>`),
		g.Group(children),
	)
}