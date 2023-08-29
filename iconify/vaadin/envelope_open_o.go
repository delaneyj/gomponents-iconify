package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeOpenO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 3.7V3h-1.5L8 0L3.4 3H2v.7L0 5v11h16V5.1l-2-1.4zM8 1.2L10.7 3H5.2L8 1.2zM3 4h10v3.7L9.5 9.4L8 8.1L6.5 9.5L3 7.8V4zM1 5.5l1-.7v2.4l-1-.4V5.5zm0 2.4l4.6 2.3l-4.6 4V7.9zm.9 7.1L8 9.7l6.1 5.3H1.9zm13.1-.8l-4.7-4.1L15 7.8v6.4zm0-7.5l-1 .5V4.9l1 .7v1.1z"/>`),
		g.Group(children),
	)
}