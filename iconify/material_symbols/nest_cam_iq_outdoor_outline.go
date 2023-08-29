package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamIqOutdoorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.075 20.25q-.275.325-.687.363t-.738-.238l-3.825-3.2q-1.25-1.05-1.8-2.55T6.75 11.55l-1.4-1.175q-.3-.275-.337-.663T5.224 9H4v2q0 .825-.588 1.413T2 13V3q.825 0 1.413.588T4 5v2h2.875l1.55-1.875q.275-.325.688-.363T9.85 5l1.375 1.15q1.525-.575 3.1-.3t2.825 1.325l3.825 3.2q.325.275.363.688t-.238.712l-7.025 8.475Zm-.9-2.075l5.775-6.925l-3.1-2.55q-1.05-.85-2.387-.825t-2.613.55l-1.525-1.25l-1.925 2.3l1.5 1.25q-.3 1.35-.075 2.688t1.275 2.212l3.075 2.55Zm0-5.5Z"/>`),
		g.Group(children),
	)
}