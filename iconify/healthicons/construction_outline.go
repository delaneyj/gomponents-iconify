package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConstructionOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.952 4h12a1 1 0 0 1 1 1v.5H32V4h4v7h-4V9.5h-1.048v.5a1 1 0 0 1-1 1h-2.666v9.253l1.19 1.19v19.843a3 3 0 0 1-3 3h-1.952a3 3 0 0 1-3-3V21.443l1.19-1.19V11h-1.762a2 2 0 0 1-2-2V6A5 5 0 0 0 13 11h-2a7 7 0 0 1 6.952-7Zm11 2h-9v3h9V6Zm-5.238 5h1.572v10.08l1.19 1.191v19.015a1 1 0 0 1-1 1h-1.952a1 1 0 0 1-1-1V22.27l1.19-1.19V11Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}