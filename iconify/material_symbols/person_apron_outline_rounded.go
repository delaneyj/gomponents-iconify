package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonApronOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12q-1.65 0-2.825-1.175T8 8q0-1.65 1.175-2.825T12 4q1.65 0 2.825 1.175T16 8q0 1.65-1.175 2.825T12 12Zm-7 8q-.425 0-.713-.288T4 19v-1.8q0-.85.425-1.563T5.6 14.55q1.5-.75 3.113-1.15T12 13q1.675 0 3.288.4t3.112 1.15q.75.375 1.175 1.088T20 17.2V19q0 .425-.287.713T19 20H5Zm7-10q.825 0 1.413-.588T14 8q0-.825-.588-1.413T12 6q-.825 0-1.413.588T10 8q0 .825.588 1.413T12 10Zm4 5.7V18h2v-.8q0-.275-.125-.5t-.375-.35q-.35-.2-.738-.363T16 15.7Zm-6-.525V16.5h4v-1.325q-.5-.1-1-.138T12 15q-.5 0-1 .038t-1 .137ZM6 18h2v-2.3q-.375.125-.763.288t-.737.362q-.25.125-.375.35T6 17.2v.8Zm10 0H8h8ZM12 8Z"/>`),
		g.Group(children),
	)
}