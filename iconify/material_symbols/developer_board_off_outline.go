package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperBoardOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.475 23.3L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM5.825 3H18q.825 0 1.413.588T20 5v2h2v2h-2v2h2v2h-2v2h2v2h-2.175L18 15.175V5H7.825l-2-2ZM11 8.175L9.825 7H11v1.175ZM12.825 10L12 9.175V7h4v3h-3.175ZM16 13.15L13.85 11H16v2.15ZM10.175 13Zm2.75-2.95ZM6 17v-4h5v4H6ZM3.2 3.175L5.025 5H4v14h14v-1.025l2 2q-.35.5-.875.763T18 21H4q-.825 0-1.413-.588T2 19V5q0-.625.338-1.113t.862-.712Zm8.8 8.8l4 4V17h-4v-5.025ZM7.025 7L11 10.975V12H6V7h1.025Z"/>`),
		g.Group(children),
	)
}