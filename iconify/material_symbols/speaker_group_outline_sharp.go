package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerGroupOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 17h9V3h-9v14Zm-2 2V1h13v18H8Zm6.5-11.5q.625 0 1.063-.438T16 6q0-.625-.438-1.063T14.5 4.5q-.625 0-1.063.438T13 6q0 .625.438 1.063T14.5 7.5Zm0 8.5q1.45 0 2.475-1.025T18 12.5q0-1.45-1.025-2.475T14.5 9q-1.45 0-2.475 1.025T11 12.5q0 1.45 1.025 2.475T14.5 16Zm0-2q-.625 0-1.063-.438T13 12.5q0-.625.438-1.063T14.5 11q.625 0 1.063.438T16 12.5q0 .625-.438 1.063T14.5 14Zm1.5 9H4V5h2v16h10v2Zm-6-6V3v14Z"/>`),
		g.Group(children),
	)
}