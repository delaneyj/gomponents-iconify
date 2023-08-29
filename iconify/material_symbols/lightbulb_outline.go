package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm-4-3v-2h8v2H8Zm.25-3q-1.725-1.025-2.738-2.75T4.5 9.5q0-3.125 2.188-5.313T12 2q3.125 0 5.313 2.188T19.5 9.5q0 2.025-1.012 3.75T15.75 16h-7.5Zm.6-2h6.3q1.125-.8 1.738-1.975T17.5 9.5q0-2.3-1.6-3.9T12 4Q9.7 4 8.1 5.6T6.5 9.5q0 1.35.613 2.525T8.85 14ZM12 14Z"/>`),
		g.Group(children),
	)
}