package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backspace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m0 10l7-7h13v14H7l-7-7zm14.41 0l2.13-2.12l-1.42-1.42L13 8.6l-2.12-2.13l-1.42 1.42L11.6 10l-2.13 2.12l1.42 1.42L13 11.4l2.12 2.13l1.42-1.42L14.4 10z"/>`),
		g.Group(children),
	)
}