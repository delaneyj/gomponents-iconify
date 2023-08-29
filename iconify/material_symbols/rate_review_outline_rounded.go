package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RateReviewOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 14h1.55q.2 0 .388-.088t.312-.212l5.6-5.55q.15-.15.15-.375t-.15-.375L12.6 5.65q-.15-.15-.375-.15t-.375.15l-5.55 5.6q-.125.125-.212.313T6 11.95v1.55q0 .2.15.35t.35.15Zm4 0H17q.425 0 .713-.288T18 13q0-.425-.288-.713T17 12h-4.5l-2 2ZM6 18l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6Zm-2-2h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}