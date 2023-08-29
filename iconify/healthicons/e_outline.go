package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 12a3 3 0 0 1 3-3h12a3 3 0 1 1 0 6h-9v6h9a3 3 0 1 1 0 6h-9v6h9a3 3 0 1 1 0 6H18a3 3 0 0 1-3-3V12Zm3-1a1 1 0 0 0-1 1v24a1 1 0 0 0 1 1h12a1 1 0 1 0 0-2H20a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h10a1 1 0 1 0 0-2H20a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h10a1 1 0 1 0 0-2H18Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}