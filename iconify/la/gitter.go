package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 2v17h2V2H6zm6 4v24h2V6h-2zm6 0v24h2V6h-2zm6 0v13h2V6h-2z"/>`),
		g.Group(children),
	)
}