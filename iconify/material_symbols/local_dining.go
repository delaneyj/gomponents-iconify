package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalDining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.375 21l-1.4-1.4l10.25-10.25q-.45-1.05-.125-2.375T14.525 4.6q1.325-1.325 2.95-1.55t2.65.8q1.025 1.025.8 2.65t-1.55 2.95q-1.05 1.1-2.375 1.425t-2.375-.125L13.375 12l7.6 7.6l-1.4 1.4l-7.6-7.55l-7.6 7.55Zm2.95-8.55l-3-3q-1.35-1.35-1.35-3.225T4.325 3l6.2 6.25l-3.2 3.2Z"/>`),
		g.Group(children),
	)
}