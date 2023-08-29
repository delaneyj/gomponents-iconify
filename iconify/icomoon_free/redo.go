package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 9a7.982 7.982 0 0 0 2.709 6l1.323-1.5a6 6 0 1 1 8.212-8.743L10.001 7h6V1l-2.343 2.343A8 8 0 0 0 .001 9z"/>`),
		g.Group(children),
	)
}