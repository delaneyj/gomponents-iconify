package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirebaseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m2.5 11.5l9-8l1 9l-5 2l-5-3Zm0 0l5-9l2 3m-7 6l1-11l3 3"/>`),
		g.Group(children),
	)
}