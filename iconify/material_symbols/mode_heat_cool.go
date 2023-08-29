package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeHeatCool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.425 16.05Q4.3 15.225 3.65 13.987T3 11.25q0-1.975.938-3.513T6 5.15q1.125-1.05 2.063-1.6L9 3v2.475q0 .625.45 1.063t1.075.437q.35 0 .65-.15t.5-.425L12 6q.95.525 1.625 1.337T14.65 9.15l-2.575 2.575l-.112-.112l-.113-.113L9 8.7l-2.8 2.8q-.55.55-.875 1.263T5 14.274q0 .5.113.938t.312.837Zm2.45-.125q-.4-.275-.637-.713T7 14.276q0-.425.163-.75t.437-.6l1.4-1.4l1.425 1.375q.075.05.125.125t.1.125l-2.775 2.775ZM7.4 22L6 20.6L19.6 7L21 8.4L17.4 12H21v2h-5.6l-.5.5l1.5 1.5H21v2h-2.6l2.1 2.1l-1.4 1.4l-2.1-2.1V22h-2v-4.6l-1.5-1.5l-.5.5V22h-2v-3.6L7.4 22Z"/>`),
		g.Group(children),
	)
}