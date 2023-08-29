package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineAddBusiness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 17h2v-3h1v-2l-1-5H2l-1 5v2h1v6h9v-6h4v3zm-6 1H4v-4h5v4zM2 4h15v2H2z"/><path fill="currentColor" d="M20 18v-3h-2v3h-3v2h3v3h2v-3h3v-2z"/>`),
		g.Group(children),
	)
}