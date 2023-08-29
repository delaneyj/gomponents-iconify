package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSomalia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#4189dd" d="M64 43c0 6.075-3.373 11-10 11H10C3.373 54 0 49.075 0 43V21c0-6.075 3.373-11 10-11h44c6.627 0 10 4.925 10 11v22"/><path fill="#e6e7e8" d="M44.939 27.04L35 27.06L31.916 17l-3.062 10.06l-9.946-.02l8.06 6.12l-3.122 9.99l8.09-6.204l8.09 6.204l-3.127-9.99z"/>`),
		g.Group(children),
	)
}