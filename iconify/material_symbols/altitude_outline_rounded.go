package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AltitudeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 12q-.425 0-.713-.288T18 11V7.8l-.925.9q-.275.275-.675.288t-.7-.288q-.275-.275-.275-.7t.275-.7l2.6-2.6q.3-.3.7-.3t.7.3l2.6 2.6q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3L20 7.825V11q0 .425-.288.713T19 12ZM3 22q-.625 0-.9-.55t.1-1.05l4-5.325q.3-.4.8-.4t.8.4L11.5 20H19l-5-6.65l-2.5 3.3L10.25 15l2.95-3.925q.3-.4.8-.4t.8.4l7 9.325q.375.5.1 1.05T21 22H3Zm8.5-2Z"/>`),
		g.Group(children),
	)
}