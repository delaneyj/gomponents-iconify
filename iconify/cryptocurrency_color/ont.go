package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ont(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#32a4be" fill-rule="nonzero"/><path fill="#fff" d="M25 24.217L9.977 9.521A8.871 8.871 0 0 1 16.2 7c4.86 0 8.8 3.854 8.8 8.609zM7 7.783l15.023 14.696A8.871 8.871 0 0 1 15.8 25C10.94 25 7 21.146 7 16.391z"/></g>`),
		g.Group(children),
	)
}