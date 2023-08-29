package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterDramaOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 20q-2.3 0-3.9-1.6T1 14.5q0-1.95 1.213-3.438T5.25 9.15q.625-2.25 2.475-3.7T12 4q2.975 0 4.988 2.063T19 11q1.875.2 2.938 1.55T23 15.475q0 1.875-1.313 3.2T18.5 20h-12Zm0-2h12q1.05 0 1.775-.725T21 15.5q0-1.05-.725-1.775T18.5 13H17v-2q0-2.075-1.463-3.538T12 6q-1.575 0-2.8.863T7.4 9.075q1.7.275 2.925 1.463T11.9 13.4q.1.425-.213.763t-.812.337q-.35 0-.613-.238T9.9 13.65q-.275-1.125-1.225-1.888T6.5 11q-1.45 0-2.475 1.025T3 14.5q0 1.45 1.025 2.475T6.5 18Z"/>`),
		g.Group(children),
	)
}