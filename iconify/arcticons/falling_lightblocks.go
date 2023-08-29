package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FallingLightblocks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 4.5h13v13H11zm0 13h13v13H11zm0 13h13v13H11zm13 0h13v13H24z"/>`),
		g.Group(children),
	)
}