package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TktonWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 18.332h39m-39 0l6.204-11.896h26.591L43.5 18.332L24 41.564L4.5 18.332zM22.799 6.436l-6.123 11.896m14.648 0L25.201 6.436M24 18.332v23.232"/>`),
		g.Group(children),
	)
}