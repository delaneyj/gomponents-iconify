package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sketchbook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.5 44.5V18.526L24 3.5l-7.5 15.026V44.5h15zM21.578 8.353h4.836"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.501 18.866a3.091 3.091 0 0 0-1.355-.313a3.386 3.386 0 0 0-3.272 3.49c0-1.928-1.287-3.49-2.873-3.49s-2.873 1.562-2.873 3.49a3.386 3.386 0 0 0-3.272-3.49a3.087 3.087 0 0 0-1.274.274m4.546 3.216V44.5m5.746-22.457V44.5"/>`),
		g.Group(children),
	)
}