package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestWakeOnPressRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 13q-.625 0-1.063-.438T19 11.5v-6q0-.625.438-1.063T20.5 4q.625 0 1.063.438T22 5.5v6q0 .625-.438 1.063T20.5 13ZM8.2 21q-.45 0-.912-.163T6.55 20.4l-3.875-4.075q-.275-.3-.275-.713t.275-.687l.15-.15q.175-.175.413-.25t.512-.025l3.25.75V4.5q0-.625.438-1.063T8.5 3q.625 0 1.063.438T10 4.5v6h.9q.2 0 .45.05t.45.15l4.1 2.05q.575.275.875.863t.2 1.212l-.625 4.45q-.125.75-.675 1.238t-1.3.487H8.2Zm-.65-2l-3.8-3.8l4.25.9V5.5q0-.225.138-.363T8.5 5q.225 0 .363.138T9 5.5v6h1.75l4.15 2.05l-.95 5.45h-6.4Z"/>`),
		g.Group(children),
	)
}