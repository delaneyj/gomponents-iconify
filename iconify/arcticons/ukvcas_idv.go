package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UkvcasIdv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 16.2V8.5h7.7m23.6 0h7.7v7.7M12.2 39.5H4.5v-7.7m39 0v7.7h-7.7M12.2 28.8H15m-2.8-9.6H15m-1.4 0v9.6m6.3 0v-9.6h1.6c2.6 0 4.8 2.2 4.8 4.8h0c0 2.6-2.2 4.8-4.8 4.8h-1.6Zm15.9-9.6l-3.1 9.6l-3.2-9.6"/>`),
		g.Group(children),
	)
}