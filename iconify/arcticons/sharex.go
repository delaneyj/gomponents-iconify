package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sharex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.343 30.878v5.35a3.642 3.642 0 0 1-2.867 3.558l-15.982 3.478a2 2 0 0 1-2.426-1.954V25.573a3.642 3.642 0 0 1 2.868-3.558l18.407-4.006l-6.107-3.294v4.623"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.657 17.122v-5.35a3.642 3.642 0 0 1 2.867-3.558l15.982-3.478a2 2 0 0 1 2.426 1.954v15.737a3.642 3.642 0 0 1-2.868 3.558l-18.407 4.006l6.107 3.294v-4.623"/>`),
		g.Group(children),
	)
}