package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalfStroke(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m320 376.4l.1-.1l26.4 14.1l85.2 45.5l-16.5-97.6l-4.8-28.7l20.7-20.5l70.1-69.3l-96.1-14.2l-29.3-4.3l-12.9-26.6l-42.8-87.8l-.1.3v289.2zm175.1 98.3c2 12-3 24.2-12.9 31.3s-23 8-33.8 2.3l-128.3-68.5l-128.3 68.5c-10.8 5.7-23.9 4.8-33.8-2.3s-14.9-19.3-12.9-31.3L169.8 329L65.6 225.9c-8.6-8.5-11.7-21.2-7.9-32.7s13.7-19.9 25.7-21.7L227 150.3L291.4 18c5.4-11 16.5-18 28.8-18s23.4 7 28.8 18l64.3 132.3l143.6 21.2c12 1.8 22 10.2 25.7 21.7s.7 24.2-7.9 32.7L470.5 329l24.6 145.7z"/>`),
		g.Group(children),
	)
}