package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsBusOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19v.5q0 .625-.438 1.063T6.5 21q-.625 0-1.063-.438T5 19.5v-1.55q-.45-.5-.725-1.113T4 15.5V6q0-2.075 1.925-3.038T12 2q4.3 0 6.15.925T20 6v9.5q0 .725-.275 1.338T19 17.95v1.55q0 .625-.438 1.063T17.5 21q-.625 0-1.063-.438T16 19.5V19H8Zm4.05-14h5.6h-11.2h5.6ZM16 12H6h12h-2ZM6 10h12V7H6v3Zm2.5 6q.625 0 1.063-.438T10 14.5q0-.625-.438-1.063T8.5 13q-.625 0-1.063.438T7 14.5q0 .625.438 1.063T8.5 16Zm7 0q.625 0 1.063-.438T17 14.5q0-.625-.438-1.063T15.5 13q-.625 0-1.063.438T14 14.5q0 .625.438 1.063T15.5 16ZM6.45 5h11.2q-.375-.425-1.613-.713T12.05 4q-2.675 0-3.913.313T6.45 5ZM8 17h8q.825 0 1.413-.588T18 15v-3H6v3q0 .825.588 1.413T8 17Z"/>`),
		g.Group(children),
	)
}