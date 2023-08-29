package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2v1h12V2h2v1h1v2h-.78l.6 3H22v2h-1v10h1v2H2v-2h1V10H2V8h1.18l.6-3H3V3h1V2h2Zm-.18 3l-.6 3h13.56l-.6-3H5.82ZM19 10h-3v10h3V10Zm-5 10V10h-4v10h4Zm-6 0V10H5v10h3Z"/>`),
		g.Group(children),
	)
}