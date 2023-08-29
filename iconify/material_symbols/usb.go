package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.825 0-1.413-.588T10 20q0-.525.275-.975T11 18.3V16H8q-.825 0-1.413-.588T6 14v-2.3q-.45-.225-.725-.675T5 10q0-.825.588-1.413T7 8q.825 0 1.413.588T9 10q0 .575-.275 1T8 11.7V14h3V6H9l3-4l3 4h-2v8h3v-2h-1V8h4v4h-1v2q0 .825-.588 1.413T16 16h-3v2.3q.475.25.738.7T14 20q0 .825-.588 1.413T12 22Z"/>`),
		g.Group(children),
	)
}