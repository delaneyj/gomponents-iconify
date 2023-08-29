package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectWithoutContact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 22v-2.25q-1.4-.45-2.35-1.6t-1.1-2.65h2q.2 1.075 1.012 1.787T17.5 18h3q.625 0 1.063.438T22 19.5V22h-6Zm3-5q-.825 0-1.413-.588T17 15q0-.825.588-1.413T19 13q.825 0 1.413.588T21 15q0 .825-.588 1.413T19 17ZM9 14q0-3.75 2.625-6.375T18 5v2q-2.925 0-4.963 2.038T11 14H9Zm4 0q0-2.075 1.463-3.538T18 9v2q-1.25 0-2.125.875T15 14h-2ZM2 11V8.5q0-.625.438-1.063T3.5 7h3q1.125 0 1.938-.713T9.45 4.5h2q-.15 1.5-1.1 2.65T8 8.75V11H2Zm3-5q-.825 0-1.413-.588T3 4q0-.825.588-1.413T5 2q.825 0 1.413.588T7 4q0 .825-.588 1.413T5 6Z"/>`),
		g.Group(children),
	)
}