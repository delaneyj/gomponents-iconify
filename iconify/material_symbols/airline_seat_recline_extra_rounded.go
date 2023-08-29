package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatReclineExtraRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 6q-.825 0-1.413-.588T6.5 4q0-.825.588-1.413T8.5 2q.825 0 1.413.588T10.5 4q0 .825-.588 1.413T8.5 6ZM13 20H7.55q-.825 0-1.512-.588T5.175 18l-1.95-9.8q-.1-.475.2-.838t.8-.362q.35 0 .625.225t.35.575L7.25 18H13q.425 0 .713.288T14 19q0 .425-.288.713T13 20Zm6 1.125L16.6 17H9.65q-.725 0-1.263-.438T7.7 15.4l-1.1-5.35q-.275-1.2.563-2.125T9.2 7q.875 0 1.588.525T11.7 8.95L12.8 14h3.25q.525 0 .975.275t.725.725l3 5.125q.2.35.088.763t-.463.612q-.35.2-.763.088T19 21.124Z"/>`),
		g.Group(children),
	)
}