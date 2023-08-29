package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.79 13.267a.75.75 0 0 0-1.086-1.034l-5.954 6.251V3.75a.75.75 0 1 0-1.5 0v14.734l-5.955-6.251a.75.75 0 1 0-1.086 1.034l7.067 7.42c.16.168.366.268.58.3a.753.753 0 0 0 .29-.001a.995.995 0 0 0 .578-.3l7.067-7.419Z"/>`),
		g.Group(children),
	)
}