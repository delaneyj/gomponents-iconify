package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmergencyShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 24q-3.025-2.575-4.512-4.775T6 15.15q0-2.55 1.738-4.35T12 9q2.525 0 4.263 1.8T18 15.15q0 1.875-1.488 4.075T12 24Zm0-7.5q.625 0 1.063-.438T13.5 15q0-.625-.438-1.063T12 13.5q-.625 0-1.063.438T10.5 15q0 .625.438 1.063T12 16.5ZM8.45 7.45l-1.4-1.4q1-1 2.275-1.525T12 4q1.4 0 2.675.525T16.95 6.05l-1.4 1.4q-.725-.725-1.637-1.088T12 6q-1 0-1.913.363T8.45 7.45Zm-2.8-2.825l-1.425-1.4Q5.8 1.65 7.8.825T12 0q2.2 0 4.2.825t3.575 2.4l-1.4 1.425Q17.1 3.375 15.45 2.687T12 2q-1.8 0-3.438.675T5.65 4.625Z"/>`),
		g.Group(children),
	)
}