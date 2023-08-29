package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.483 2.959L22 20.809V2.959ZM2 2.959V21.04L9.425 2.96Zm7.069 14.324h2.784l1.748 3.433h2.537l-4.1-10.843Z"/>`),
		g.Group(children),
	)
}