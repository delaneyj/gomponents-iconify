package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m6.906 6.594l-.718.687l-3.907 3.907l-.687.718L16 26.312l14.406-14.406l-.687-.719l-3.907-3.906l-.718-.687L16 15.687zm-.031 2.843l8.406 8.376l.719.687l.719-.688l8.406-8.375l2.438 2.438L16 23.469L4.437 11.875z"/>`),
		g.Group(children),
	)
}