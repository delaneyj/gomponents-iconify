package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Columns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 6.5v11c0 .465 0 .697.038.89a2 2 0 0 0 1.572 1.571c.193.039.425.039.89.039s.697 0 .89-.039a2 2 0 0 0 1.571-1.57c.039-.194.039-.426.039-.891v-11c0-.465 0-.697-.039-.89a2 2 0 0 0-1.57-1.572C17.196 4 16.964 4 16.5 4s-.697 0-.89.038a2 2 0 0 0-1.572 1.572C14 5.803 14 6.035 14 6.5Zm-9 0v11c0 .465 0 .697.038.89a2 2 0 0 0 1.572 1.571c.193.039.425.039.89.039s.697 0 .89-.039a2 2 0 0 0 1.571-1.57c.039-.194.039-.426.039-.891v-11c0-.465 0-.697-.039-.89a2 2 0 0 0-1.57-1.572C8.196 4 7.964 4 7.5 4s-.697 0-.89.038A2 2 0 0 0 5.038 5.61C5 5.803 5 6.035 5 6.5Z"/>`),
		g.Group(children),
	)
}