package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitlife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.053 21.917C23.473 13.035 37.98 8.84 41.73 12.64c3.75 3.8-7.535 16.85-13.845 12.252"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.053 21.917l-1.711 4.044c-1.216 3.563-3.872 2.301-5.686 1.666c-2.953-1.035-5.414-.696-6.566 2.45c-.752 2.055-1.35 3.883-1.76 4.83c-.41.946-2.34 1.216-3.805.998c1.28.481 2.593 1.04 3.567.786c1.151-.3.992-.516 2.974-4.735c1.199-2.936 2.302-2.68 4.306-1.617c4.562 2.42 7.655 1.426 9.206-1.285c1.038-1.814 1.285-3.826 2.308-4.163"/>`),
		g.Group(children),
	)
}