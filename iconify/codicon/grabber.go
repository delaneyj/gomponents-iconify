package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grabber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 6H1v1h14V6zm0 3H1v1h14V9z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}