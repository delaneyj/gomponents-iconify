package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 7.33V2.5h4.618l1.453 2h-4.54a1 1 0 0 0-.965.739L1.5 7.329ZM.005 13.087A.758.758 0 0 1 0 13V2.12C0 1.501.501 1 1.12 1h5.191c.359 0 .696.172.907.462L9.062 4h3.954c.409 0 .772.196 1 .5h.95a1 1 0 0 1 .966 1.261l-2.029 7.5a1 1 0 0 1-.965.739H1.002a1 1 0 0 1-.997-.913Zm1.65-.587L3.414 6h10.9l-1.759 6.5h-10.9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}