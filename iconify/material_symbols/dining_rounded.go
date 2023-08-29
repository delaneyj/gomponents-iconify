package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiningRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.75 19q.325 0 .537-.213t.213-.537v-6q.65-.2 1.075-.713T11 10.3V6.5q0-.2-.15-.35T10.5 6q-.2 0-.35.15T10 6.5V9h-.75V6.5q0-.2-.15-.35T8.75 6q-.2 0-.35.15t-.15.35V9H7.5V6.5q0-.2-.15-.35T7 6q-.2 0-.35.15t-.15.35v3.8q0 .725.425 1.238T8 12.25v6q0 .325.213.537T8.75 19Zm6 0q.325 0 .537-.213t.213-.537v-5.6q.825-.4 1.288-1.275t.462-2.05q0-1.425-.713-2.375T14.75 6q-1.075 0-1.788.95t-.712 2.375q0 1.175.463 2.05T14 12.65v5.6q0 .325.213.538t.537.212ZM4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v16q0 .825-.588 1.413T20 22H4Z"/>`),
		g.Group(children),
	)
}