package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentBulletList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.031.031v3.938h3.871L10.031.031z"/><path d="M9 5V0H2.042v16H14V5H9zM5 8H4V7h1v1zm0 2H4V9h1v1zm-.062 2h-1v-1h1v1zm1.02 0v-1h5v1h-5zm5-2h-5V9h5v1zm0-2h-5V7h5v1z"/></g>`),
		g.Group(children),
	)
}