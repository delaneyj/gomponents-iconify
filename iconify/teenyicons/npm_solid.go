package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NpmSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 4h15v7H7v2H4v-2H0V4Zm4 6V5H1v5h1V6h1v4h1Zm1-5v7h1v-2h2V5H5Zm4 0v5h1V6h1v4h1V6h1v4h1V5H9ZM6 9V6h1v3H6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}