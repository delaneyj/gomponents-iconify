package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoopInfinity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.269 28.202l-1.517 1.709a8.352 8.352 0 1 1 0-11.822l10.496 11.822a8.352 8.352 0 1 0 0-11.822l-1.517 1.709"/>`),
		g.Group(children),
	)
}