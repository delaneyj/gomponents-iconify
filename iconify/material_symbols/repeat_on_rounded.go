package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatOnRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23q-.825 0-1.413-.588T1 21V3q0-.825.588-1.413T3 1h18q.825 0 1.413.588T23 3v18q0 .825-.588 1.413T21 23H3Zm14-6H6.85l.85-.85q.3-.3.3-.713t-.3-.712q-.3-.3-.713-.3t-.712.3L3.7 17.3q-.275.275-.275.7t.275.7l2.6 2.6q.275.275.688.275t.712-.3q.3-.3.288-.725t-.313-.725L6.85 19H18q.425 0 .713-.287T19 18v-4.025q0-.425-.288-.7T18 13q-.425 0-.713.288T17 14v3ZM7 7h10.15l-.85.85q-.3.3-.3.713t.3.712q.3.3.712.3t.713-.3L20.3 6.7q.275-.275.275-.7t-.275-.7l-2.6-2.6q-.275-.275-.688-.275t-.712.3q-.3.3-.288.725t.313.725L17.15 5H6q-.425 0-.713.288T5 6v4.025q0 .425.288.7T6 11q.425 0 .713-.288T7 10V7Z"/>`),
		g.Group(children),
	)
}