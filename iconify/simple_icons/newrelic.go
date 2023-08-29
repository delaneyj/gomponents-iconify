package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newrelic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.002 14.31v7.383L12 24V12L1.608 6v4.616ZM12 0L2.823 5.298l3.998 2.308L12 4.616l6.393 3.692v7.384l-5.178 2.99V23.3l9.176-5.3V6Z"/>`),
		g.Group(children),
	)
}