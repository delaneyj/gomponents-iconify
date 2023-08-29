package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 76h-85.33a4 4 0 0 1-2.4-.8l-27.74-20.8a12.05 12.05 0 0 0-7.2-2.4H40a12 12 0 0 0-12 12v136a12 12 0 0 0 12 12h176.89A11.12 11.12 0 0 0 228 200.89V88a12 12 0 0 0-12-12Zm4 124.89a3.12 3.12 0 0 1-3.11 3.11H40a4 4 0 0 1-4-4V64a4 4 0 0 1 4-4h53.33a4 4 0 0 1 2.4.8l27.74 20.8a12.05 12.05 0 0 0 7.2 2.4H216a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}