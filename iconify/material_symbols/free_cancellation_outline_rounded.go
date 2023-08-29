package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreeCancellationOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.525 19.675l3.55-3.55q.3-.3.7-.288t.7.313q.275.3.275.7t-.275.7L17.25 21.8q-.3.3-.7.3t-.7-.3l-2.15-2.15q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.425 1.425ZM9 15.4l-.9.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l.9-.9l-.9-.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l.9.9l.9-.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-.9.9l.9.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-.9-.9ZM5 22q-.825 0-1.413-.588T3 20V6q0-.825.588-1.413T5 4h1V3q0-.425.288-.713T7 2q.425 0 .713.288T8 3v1h8V3q0-.425.288-.713T17 2q.425 0 .713.288T18 3v1h1q.825 0 1.413.588T21 6v6.35l-2 2.025V10H5v10h6.25l1.975 2H5ZM5 8h14V6H5v2Zm0 0V6v2Z"/>`),
		g.Group(children),
	)
}