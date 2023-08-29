package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LowPriority(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 11.475q0 1.775 1.188 3.05T8.15 15.95L6.6 14.4L8 13l4 4l-4 4l-1.4-1.4L8.2 18q-2.625-.15-4.413-2.025T2 11.5q0-2.725 1.888-4.612T8.5 5H12v2H8.5Q6.625 7 5.312 8.3T4 11.475ZM14 18v-2h8v2h-8Zm0-5.5v-2h8v2h-8ZM14 7V5h8v2h-8Z"/>`),
		g.Group(children),
	)
}