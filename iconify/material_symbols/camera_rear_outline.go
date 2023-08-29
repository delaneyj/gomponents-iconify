package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRearOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17V4q0-.825.588-1.413T7 2h10q.825 0 1.413.588T19 4v13h-2V4H7v13H5Zm7-7q.825 0 1.413-.588T14 8q0-.825-.588-1.413T12 6q-.825 0-1.413.588T10 8q0 .825.588 1.413T12 10ZM9.85 22.5l-1.4-1.4l1.1-1.1H5v-2h4.55l-1.1-1.1l1.4-1.4l3.5 3.5l-3.5 3.5ZM14 20v-2h5v2h-5Zm-2-9.5Z"/>`),
		g.Group(children),
	)
}