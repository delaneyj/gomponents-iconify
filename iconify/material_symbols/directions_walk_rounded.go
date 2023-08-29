package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsWalkRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 5.5q-.825 0-1.413-.588T11.5 3.5q0-.825.588-1.413T13.5 1.5q.825 0 1.413.588T15.5 3.5q0 .825-.588 1.413T13.5 5.5ZM8.325 23q-.55 0-.875-.388t-.2-.912L9.8 8.9L8 9.6V12q0 .425-.288.712T7 13q-.425 0-.713-.288T6 12V9.625q0-.6.325-1.1t.9-.75L11.05 6.15q.725-.3 1.475-.063T13.7 7l1 1.6q.55.875 1.425 1.5t2 .825q.375.05.625.337t.25.638q0 .475-.338.787t-.762.238q-1.325-.2-2.462-.838T13.5 10.5l-.6 3l1.475 1.4q.3.3.463.663t.162.787V22q0 .425-.288.713T14 23q-.425 0-.713-.288T13 22v-5l-2.1-2l-1.625 7.25q-.075.3-.338.525T8.325 23Z"/>`),
		g.Group(children),
	)
}