package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsRemoteOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 23V9h8v14H8Zm4-8.75q.525 0 .888-.363T13.25 13q0-.525-.363-.888T12 11.75q-.525 0-.888.363T10.75 13q0 .525.363.888t.887.362Zm-3.55-6.8l-1.4-1.4Q8 5.1 9.263 4.55T12 4q1.475 0 2.738.55t2.212 1.5l-1.4 1.4q-.675-.675-1.588-1.063T12 6q-1.05 0-1.963.388T8.45 7.45Zm-2.8-2.8L4.2 3.2Q5.7 1.725 7.7.863T12 0q2.3 0 4.288.875T19.75 3.25l-1.4 1.4q-1.2-1.25-2.838-1.95T12 2q-1.875 0-3.513.7T5.65 4.65ZM10 21h4V11h-4v10Zm0 0h4h-4Z"/>`),
		g.Group(children),
	)
}