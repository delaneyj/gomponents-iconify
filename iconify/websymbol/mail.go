package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M105 5h1152q43 0 74 31t31 74v785q0 44-31 74.5t-74 30.5H105q-43 0-74-30.5T0 895V110q0-43 31-74t74-31zm69 157l507 423l507-423H174zm-17 681h1048V285L711 697h-61L157 285v558z"/>`),
		g.Group(children),
	)
}