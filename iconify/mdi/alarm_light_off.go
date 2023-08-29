package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmLightOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 14.8l-9-9c.9-.5 1.9-.8 3-.8c3.3 0 6 2.7 6 6v3.8m2.1-10l-1.4-1.4l-2.1 2.1L18 6.9l2.1-2.1m-.6 5.7v2h3v-2h-3m-15 0h-3v2h3v-2M1.1 3l5.5 5.5c-.4.7-.6 1.6-.6 2.5v8h11.1l1 1H6c-1.1 0-2 .9-2 2h16.1l.7.7l1.3-1.3L2.4 1.7L1.1 3M13 1h-2v3h2V1Z"/>`),
		g.Group(children),
	)
}