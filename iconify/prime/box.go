package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.73 16.52V7.59a.73.73 0 0 0-.08-.33a.74.74 0 0 0-.36-.36l-8-3.58a.75.75 0 0 0-.62 0l-8 3.58a.8.8 0 0 0-.44.69v8.82a.83.83 0 0 0 .44.69l8 3.58a.72.72 0 0 0 .62 0l8-3.58a.77.77 0 0 0 .44-.58Zm-16-7.78l6.5 2.92v7.18l-6.5-2.91Zm8 2.92l6.5-2.92v7.19l-6.5 2.91ZM12 4.82l6.17 2.77L12 10.35L5.83 7.59Z"/>`),
		g.Group(children),
	)
}