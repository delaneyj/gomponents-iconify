package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CleaningBucketOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.75 23q-.75 0-1.313-.488t-.662-1.237L5 8h14l-1.775 13.275q-.1.75-.663 1.238T15.25 23h-6.5Zm0-2h6.5l1.45-11H7.275L8.75 21ZM12 16q1.25 0 2.125-.875T15 13v-2h-2v2q0 .425-.288.713T12 14q-.425 0-.713-.288T11 13v-2H9v2q0 1.25.875 2.125T12 16Zm3-9q-.625 0-1.063-.438T13.5 5.5q0-.625.438-1.063T15 4q.625 0 1.063.438T16.5 5.5q0 .625-.438 1.063T15 7Zm-5-1q-1.05 0-1.775-.725T7.5 3.5q0-1.05.725-1.775T10 1q1.05 0 1.775.725T12.5 3.5q0 1.05-.725 1.775T10 6Zm5.25 15h-6.5h6.5Z"/>`),
		g.Group(children),
	)
}