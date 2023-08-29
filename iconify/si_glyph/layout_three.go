package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M17.021 16.979h-16V1h16v15.979zM2 16h14V2H2v14z"/><path d="M3 10v4.953h4.977V10H3zm7 0v4.992h5.002V10H10zM3 3v4.96h5.01V3H3zm7 0v4.963h5V3h-5z"/></g>`),
		g.Group(children),
	)
}