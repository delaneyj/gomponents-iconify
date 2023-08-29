package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#1fc5c9"/><path fill="#fff" d="m11.296 22.472l2.164-2.164a5 5 0 0 0 6.848-6.848l2.164-2.164a8 8 0 0 1-11.176 11.176zm-1.768-1.768A8 8 0 0 1 20.704 9.528l-2.164 2.164a5 5 0 0 0-6.848 6.848zM24 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4zM8 26a2 2 0 1 1 0-4a2 2 0 0 1 0 4z"/></g>`),
		g.Group(children),
	)
}