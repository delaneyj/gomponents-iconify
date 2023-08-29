package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetailsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.7 21q-.575 0-.863-.488t-.012-.987l8.3-14.95q.275-.5.875-.5t.875.5l8.3 14.95q.275.5-.013.988T20.3 21H3.7Zm1.7-2H11V8.925L5.4 19Zm7.6 0h5.6L13 8.925V19Z"/>`),
		g.Group(children),
	)
}