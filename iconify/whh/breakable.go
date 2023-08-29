package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Breakable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 698v198h128q26 0 45 18.5t19 45t-19 45.5t-45 19H192q-27 0-45.5-19T128 959.5t18.5-45T192 896h128V698Q183 672 91.5 554.5T0 282Q0 121 98 0h222l64 192l-192 64l192 320l-64-256l192-64L448 0h221q99 121 99 282q0 155-91.5 272.5T448 698z"/>`),
		g.Group(children),
	)
}