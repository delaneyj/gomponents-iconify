package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdfcFirstBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h37v12.333h-37zm0 24.667h12.333V42.5H5.5zm0-12.334h24.667v12.333H5.5z"/>`),
		g.Group(children),
	)
}