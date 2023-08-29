package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gbp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#bc3fe0"/><path fill="#fff" d="M11.087 14.815v-2.332c0-3.676 2.219-5.983 6.075-5.983c2.932 0 4.57 1.242 5.838 2.84l-2.483 1.9c-.951-1.165-1.85-1.85-3.328-1.85c-1.77 0-2.827 1.217-2.827 3.17v2.255h6.578v2.637h-6.578v4.335h8.585V24.5H9v-1.977l2.087-.609v-4.462H9v-2.637z"/></g>`),
		g.Group(children),
	)
}