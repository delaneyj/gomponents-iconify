package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Realtor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.1 6L5.5 25H11v17h26V25h5.5L24.1 6Zm-2.754 14.413v15m0-15h-2.183m-.647 15h5.66"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.346 26.073a5.66 5.66 0 0 1 5.66-5.66h0c1.438 0 2.028.287 2.478 1.36"/>`),
		g.Group(children),
	)
}