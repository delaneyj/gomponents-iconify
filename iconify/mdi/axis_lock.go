package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AxisLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.8 5c.6 0 1.2.6 1.2 1.3v3.5c0 .6-.6 1.2-1.3 1.2h-5.5c-.6 0-1.2-.6-1.2-1.3V6.2c0-.6.6-1.2 1.2-1.2V3.5C16.2 2.1 17.6 1 19 1s2.8 1.1 2.8 2.5V5m-1.3 0V3.5c0-.8-.7-1.3-1.5-1.3s-1.5.5-1.5 1.3V5h3M11 13.82l9.39 5.43l-1 1.75L10 15.56L3.17 19.5l-1-1.73L9 13.82V3h2v10.82Z"/>`),
		g.Group(children),
	)
}