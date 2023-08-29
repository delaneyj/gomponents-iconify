package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viewport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 4H0V0h4v1H1zm11-3V0h4v4h-1V1zm3 11h1v4h-4v-1h3zM4 15v1H0v-4h1v3z"/><path fill="currentColor" d="M13 3v10H3V3h10zm1-1H2v12h12V2z"/>`),
		g.Group(children),
	)
}