package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentViewGallerySixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2H4.5A2.5 2.5 0 0 0 2 4.5v7A2.5 2.5 0 0 0 4.5 14H10V2Zm1 12h.5a2.5 2.5 0 0 0 2.5-2.5V11h-3v3Zm3-4V6h-3v4h3Zm0-5v-.5A2.5 2.5 0 0 0 11.5 2H11v3h3ZM4 4.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-3ZM5 5v2h2V5H5Zm-.5 4h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1 0-1Zm0 2h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}