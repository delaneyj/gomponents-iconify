package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertTemplate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 3h2v1H6zm3 0h2v1H9zm5 0v4h-3V6h2V4h-1V3zM5 6h2v1H5zm3 0h2v1H8zM3 4v2h1v1H2V3h3v1zm3 5h2v1H6zm3 0h2v1H9zm5 0v4h-3v-1h2v-2h-1V9zm-9 3h2v1H5zm3 0h2v1H8zm-5-2v2h1v1H2V9h3v1zm12-9H1v14h14V1zm1-1v16H0V0h16z"/>`),
		g.Group(children),
	)
}