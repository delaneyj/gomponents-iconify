package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pulse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M11.5 8L8.8 5.4L6.6 8.5L5.5 1.6L2.38 8H0v2h3.6l.9-1.8l.9 5.4L9 8.5l1.6 1.5H14V8h-2.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}