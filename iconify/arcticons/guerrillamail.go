package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Guerrillamail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.8 16.2h20.8v19.5a4.388 4.388 0 0 1-3.9 3.9h-13a4.388 4.388 0 0 1-3.9-3.9Zm20.8 0h1.3V11h-5.85l-2.6-2.6h-6.5l-2.6 2.6H4.5v5.2h1.3m24.7 1.95h13m-13 7.8h11.05m-11.05 7.8h7.15"/>`),
		g.Group(children),
	)
}