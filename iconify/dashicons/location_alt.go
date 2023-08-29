package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m13 13.14l1.17-5.94c.79-.43 1.33-1.25 1.33-2.2a2.5 2.5 0 0 0-5 0c0 .95.54 1.77 1.33 2.2zm0-9.64c.83 0 1.5.67 1.5 1.5s-.67 1.5-1.5 1.5s-1.5-.67-1.5-1.5s.67-1.5 1.5-1.5zm1.72 4.8L18 6.97v9L13.12 18L7 15.97l-5 2v-9l5-2l4.27 1.41l1.73 7.3z"/>`),
		g.Group(children),
	)
}