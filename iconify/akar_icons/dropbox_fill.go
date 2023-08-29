package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropboxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 2l6 3.75L6 9.5L0 5.75L6 2Zm12 0l6 3.75l-6 3.75l-6-3.75L18 2ZM0 13.25L6 9.5l6 3.75L6 17l-6-3.75ZM18 9.5l6 3.75L18 17l-6-3.75l6-3.75ZM6 18.25l6-3.75l6 3.75L12 22l-6-3.75Z"/>`),
		g.Group(children),
	)
}