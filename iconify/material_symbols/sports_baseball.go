package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsBaseball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.175 18.225q-1.05-1.35-1.613-2.938T2 12q0-1.7.563-3.288t1.612-2.937Q5.625 6.9 6.45 8.538T7.275 12q0 1.825-.825 3.463t-2.275 2.762ZM12 22q-1.8 0-3.438-.6t-2.987-1.75q1.725-1.425 2.7-3.412T9.25 12q0-2.25-.975-4.237t-2.7-3.413Q6.925 3.2 8.563 2.6T12 2q1.8 0 3.438.6t2.987 1.75q-1.725 1.425-2.7 3.413T14.75 12q0 2.25.975 4.237t2.7 3.413q-1.35 1.15-2.988 1.75T12 22Zm7.825-3.775q-1.45-1.125-2.275-2.763T16.725 12q0-1.825.825-3.463t2.275-2.762q1.05 1.35 1.613 2.938T22 12q0 1.7-.563 3.288t-1.612 2.937Z"/>`),
		g.Group(children),
	)
}