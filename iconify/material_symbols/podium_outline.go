package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodiumOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 4.5q0 .825-.588 1.413T10.5 6.5q-.325 0-.6-.088t-.55-.287q-.6.2-.963.725T8.025 8H21l-1 7h-4.9v-2h3.175q.125-.75.213-1.5T18.7 10H5.3q.125.75.213 1.5t.212 1.5H8.9v2H4L3 8h3q0-1.225.675-2.225T8.5 4.3q.075-.775.65-1.287T10.5 2.5q.825 0 1.413.588T12.5 4.5ZM9.775 19h4.45l.575-6H9.2l.575 6ZM8 21l-.75-7.8q-.1-.875.5-1.538T9.225 11h5.55q.875 0 1.475.663t.5 1.537L16 21H8Z"/>`),
		g.Group(children),
	)
}