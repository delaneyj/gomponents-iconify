package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.5 19h11c.465 0 .697 0 .89-.039a2 2 0 0 0 1.571-1.57c.039-.194.039-.426.039-.891s0-.697-.039-.89a2 2 0 0 0-1.57-1.572C18.196 14 17.964 14 17.5 14h-11c-.465 0-.697 0-.89.038a2 2 0 0 0-1.572 1.572C4 15.803 4 16.035 4 16.5s0 .697.038.89a2 2 0 0 0 1.572 1.571c.193.039.425.039.89.039Zm0-9h11c.465 0 .697 0 .89-.039a2 2 0 0 0 1.571-1.57C20 8.196 20 7.964 20 7.5s0-.697-.039-.89a2 2 0 0 0-1.57-1.572C18.196 5 17.964 5 17.5 5h-11c-.465 0-.697 0-.89.038A2 2 0 0 0 4.038 6.61C4 6.803 4 7.035 4 7.5s0 .697.038.89A2 2 0 0 0 5.61 9.961c.193.039.425.039.89.039Z"/>`),
		g.Group(children),
	)
}