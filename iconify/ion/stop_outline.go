package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="320" height="320" x="96" y="96" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="24" ry="24"/>`),
		g.Group(children),
	)
}