package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 16C2 10 7.234 0 16 0c-4.109 3.297-6 11-9 11H4l-3 5H0z"/>`),
		g.Group(children),
	)
}