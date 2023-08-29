package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm8.146-3.354a.5.5 0 0 1 .708 0l2 2a.5.5 0 0 1 0 .708L6.707 11H4.5a.5.5 0 0 1-.5-.5V8.293l4.146-4.147Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}