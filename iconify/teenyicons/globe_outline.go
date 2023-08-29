package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 15v-3M4 14.5h7M11.469 1A6 6 0 1 1 3.5 9.972m4-7.472a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/>`),
		g.Group(children),
	)
}