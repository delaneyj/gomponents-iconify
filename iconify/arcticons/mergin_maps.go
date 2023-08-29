package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MerginMaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5v6.6m0 23.8v6.6m0-28v19m9.5-9.5a9.5 9.5 0 0 1-9.5 9.5h0a9.5 9.5 0 0 1-9.5-9.5h0a9.5 9.5 0 0 1 9.5-9.5h0a9.5 9.5 0 0 1 9.5 9.5h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.1 12.1h23.8v23.8H12.1z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5h37v37h-37z"/>`),
		g.Group(children),
	)
}