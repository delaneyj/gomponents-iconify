package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalLibraryRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17.15V10q0-.8.588-1.35t1.387-.5q1.975.3 3.763 1.163T12 11.55q1.475-1.375 3.263-2.238t3.762-1.162q.8-.05 1.388.5T21 10v7.15q0 .8-.525 1.363t-1.325.612q-1.6.25-3.1.825t-2.8 1.525q-.275.225-.588.337t-.662.113q-.35 0-.663-.113t-.587-.337q-1.3-.95-2.8-1.525t-3.1-.825q-.8-.05-1.325-.612T3 17.15ZM12 9q-1.65 0-2.825-1.175T8 5q0-1.65 1.175-2.825T12 1q1.65 0 2.825 1.175T16 5q0 1.65-1.175 2.825T12 9Z"/>`),
		g.Group(children),
	)
}