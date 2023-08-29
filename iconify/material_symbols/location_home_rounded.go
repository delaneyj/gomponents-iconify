package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationHomeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15q1.25 0 2.125-.875T15 12q0-1.25-.875-2.125T12 9q-1.25 0-2.125.875T9 12q0 1.25.875 2.125T12 15Zm0 1q-1.45 0-2.813.413T6.6 17.6q-.3.175-.45.475t-.15.65V19h12v-.275q0-.35-.15-.65t-.45-.475q-1.225-.775-2.587-1.188T12 16Zm-6 5q-.825 0-1.413-.588T4 19v-9q0-.475.213-.9t.587-.7l6-4.5q.525-.4 1.2-.4t1.2.4l6 4.5q.375.275.588.7T20 10v9q0 .825-.588 1.413T18 21H6Z"/>`),
		g.Group(children),
	)
}