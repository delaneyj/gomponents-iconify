package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.883 10.577a2.217 2.217 0 0 1 1.796-2.186l19.673-3.848a2.217 2.217 0 0 1 2.606 1.743m-24.075 9.55a2.217 2.217 0 0 1 1.796-2.186l19.673-3.848A2.217 2.217 0 0 1 41 11.977m0 0l-.031 19.936M16.883 15.836v22.717"/><circle cx="11.946" cy="38.553" r="4.947" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.023" cy="31.913" r="4.947" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}