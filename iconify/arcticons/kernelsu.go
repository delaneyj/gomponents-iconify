package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kernelsu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h37v37h-37z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26 9h13v13H26zM9 26h13v13H9z"/>`),
		g.Group(children),
	)
}