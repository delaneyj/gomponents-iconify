package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CicsRegionAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M31 13v-2h-4c-1.103 0-2 .897-2 2v2c0 1.103.897 2 2 2h2v2h-4v2h4c1.103 0 2-.897 2-2v-2c0-1.103-.897-2-2-2h-2v-2h4zm-14 0v6c0 1.103.897 2 2 2h4v-2h-4v-6h4v-2h-4c-1.103 0-2 .897-2 2zm-8 0h2v6H9v2h6v-2h-2v-6h2v-2H9v2zm-8 0v6c0 1.103.897 2 2 2h4v-2H3v-6h4v-2H3c-1.103 0-2 .897-2 2z"/>`),
		g.Group(children),
	)
}