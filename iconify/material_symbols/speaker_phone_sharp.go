package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerPhoneSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.45 8.5L7 7.05q1.025-.975 2.313-1.513T12 5q1.4 0 2.688.537T17 7.05L15.55 8.5q-.725-.725-1.625-1.113T12 7q-1.025 0-1.925.388T8.45 8.5Zm-2.8-2.9l-1.4-1.4q1.575-1.55 3.563-2.375T12 1q2.2 0 4.188.825T19.75 4.2l-1.4 1.4q-1.275-1.25-2.925-1.925T12 3q-1.775 0-3.425.675T5.65 5.6ZM8 22V10h8v12H8Z"/>`),
		g.Group(children),
	)
}