package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Streetview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 11q-2.075 0-3.538-1.463T13 6q0-2.075 1.463-3.538T18 1q2.075 0 3.538 1.463T23 6q0 2.075-1.463 3.538T18 11ZM3.6 20.4q-.275-.275-.438-.625T3 19V5q0-.825.588-1.413T5 3h6.7q-.35.675-.525 1.438T11 6q0 1.475.55 2.738t1.5 2.212L3.6 20.4Zm8.4.6v-5.4q0-1.05.638-1.887T14.3 12.65q.875-.2 1.8-.3t1.9-.1q.8 0 1.538.063T21 12.5V19q0 .825-.588 1.413T19 21h-7Z"/>`),
		g.Group(children),
	)
}