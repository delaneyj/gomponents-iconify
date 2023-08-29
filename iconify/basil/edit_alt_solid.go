package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.607 3.5a.5.5 0 0 1 .353.146l2.829 2.829a.5.5 0 0 1 0 .707l-9.193 9.192a.5.5 0 0 1-.227.13l-3.828 1a.5.5 0 0 1-.61-.61l1-3.828a.5.5 0 0 1 .13-.227l9.192-9.193a.5.5 0 0 1 .354-.146ZM4 19.25a.75.75 0 0 0 0 1.5h15a.75.75 0 0 0 0-1.5H4Z"/>`),
		g.Group(children),
	)
}