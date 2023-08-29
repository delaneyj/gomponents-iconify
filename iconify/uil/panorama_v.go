package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.27 12a21.11 21.11 0 0 1 1.64-8.58a1 1 0 0 0-.07-1A1 1 0 0 0 18 2H6a1 1 0 0 0-.84.46a1 1 0 0 0-.07 1A21.11 21.11 0 0 1 6.73 12a21.11 21.11 0 0 1-1.64 8.58a1 1 0 0 0 .07 1A1 1 0 0 0 6 22h12a1 1 0 0 0 .84-.46a1 1 0 0 0 .07-1A21.11 21.11 0 0 1 17.27 12Zm-.75 8h-9a24.77 24.77 0 0 0 1.25-8a24.77 24.77 0 0 0-1.29-8h9a24.77 24.77 0 0 0-1.25 8a24.77 24.77 0 0 0 1.29 8Z"/>`),
		g.Group(children),
	)
}