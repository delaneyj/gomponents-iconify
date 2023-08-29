package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Satellite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M26.736 22.723a3.2 3.2 0 1 1-4.224-1.625a3.201 3.201 0 0 1 4.224 1.625Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.282 4.5l7.812 13m7.812 13l7.812 13"/>`),
		g.Group(children),
	)
}