package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdminMedsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12Zm-3.8 3.775q.725.725 1.713.738t1.712-.713l1.4-1.4L9.6 10.975l-1.4 1.4q-.725.725-.725 1.7t.725 1.7Zm7.6-7.55q-.725-.7-1.713-.725t-1.712.7L11 9.575L14.425 13l1.375-1.375q.725-.725.713-1.7t-.713-1.7ZM3 21V3h6.2q.325-.9 1.088-1.45T12 1q.95 0 1.713.55T14.8 3H21v18H3Zm9-16.75q.325 0 .537-.213t.213-.537q0-.325-.213-.537T12 2.75q-.325 0-.537.213t-.213.537q0 .325.213.537T12 4.25Z"/>`),
		g.Group(children),
	)
}