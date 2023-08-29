package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatchJobStep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 20v4h-4v4h-4v2h6v-4h4v-4h4v-2h-6zM8 4h8v6c0 1.103.897 2 2 2h6v4h2v-6a1 1 0 0 0-.293-.707l-7-7A1 1 0 0 0 18 2H8c-1.103 0-2 .897-2 2v24c0 1.103.897 2 2 2h4v-2H8V4zm15.586 6H18V4.414L23.586 10z"/>`),
		g.Group(children),
	)
}