package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashcube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24.5 3l-.313.281L20.47 7H11c-2.75 0-5 2.25-5 5v10c0 2.75 2.25 5 5 5h10c2.75 0 5-2.25 5-5V3zM24 6.313V22c0 1.668-1.332 3-3 3H11c-1.668 0-3-1.332-3-3V12c0-1.668 1.332-3 3-3h10.313zM13 12c-1.094 0-2 .906-2 2v6c0 1.094.906 2 2 2h10.406l-1.687-1.719l-.719-.718V14c0-1.094-.906-2-2-2zm0 2h6v6h-6z"/>`),
		g.Group(children),
	)
}