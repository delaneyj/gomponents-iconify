package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenRotationUpRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.175 20H8q-.825 0-1.413-.588T6 18V7.825l2 2V18h6.175l-1.125-1.125q-.275-.275-.275-.688t.275-.712q.3-.3.713-.3t.712.3L17.3 18.3q.3.3.3.7t-.3.7l-2.85 2.85q-.3.3-.7.288t-.7-.313q-.275-.3-.288-.7t.288-.7L14.175 20ZM18 16.175l-2-2V6H9.825l1.125 1.125q.275.275.275.688t-.275.712q-.3.3-.712.3t-.713-.3L6.7 5.7q-.3-.3-.3-.7t.3-.7l2.85-2.85q.3-.3.7-.287t.7.312q.275.3.288.7t-.288.7L9.825 4H16q.825 0 1.413.588T18 6v10.175Z"/>`),
		g.Group(children),
	)
}