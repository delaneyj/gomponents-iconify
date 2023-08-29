package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOfDavid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M58 17.448H41.015L32 2l-9.014 15.448H6L14.492 32L6 46.553h16.986L32 62l9.015-15.447H58L49.508 32L58 17.448zm-6.146 3.533L47.465 28.5l-4.388-7.519h8.777zM32 9.001l4.93 8.447h-9.86L32 9.001zm-19.853 11.98h8.776L16.536 28.5l-4.389-7.519zm0 22.039l4.389-7.52l4.388 7.52h-8.777zM32 54.998l-4.93-8.445h9.859L32 54.998zm6.992-11.978H25.009L18.578 32l6.431-11.02h13.983L45.422 32l-6.43 11.02zm12.862 0h-8.776l4.388-7.52l4.388 7.52z"/>`),
		g.Group(children),
	)
}