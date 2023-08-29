package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOnOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 22q-.825 0-1.413-.588T8 20v-9L6 8V4q0-.825.588-1.413T8 2h8q.825 0 1.413.588T18 4v4l-2 3v9q0 .825-.588 1.413T14 22h-4Zm2-6.5q-.625 0-1.063-.438T10.5 14q0-.625.438-1.063T12 12.5q.625 0 1.063.438T13.5 14q0 .625-.438 1.063T12 15.5ZM8 5h8V4H8v1Zm8 2H8v.4l2 3V20h4v-9.6l2-3V7Zm-4 5Z"/>`),
		g.Group(children),
	)
}