package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11a7 7 0 1 0-14 0m11 3.5v2c0 .465 0 .697.038.89a2 2 0 0 0 1.572 1.572c.193.038.425.038.89.038s.697 0 .89-.038a2 2 0 0 0 1.571-1.572c.039-.193.039-.425.039-.89v-2c0-.465 0-.697-.039-.89a2 2 0 0 0-1.57-1.572C19.196 12 18.964 12 18.5 12s-.697 0-.89.038a2 2 0 0 0-1.572 1.571c-.038.194-.038.426-.038.891Zm-8 0v2c0 .465 0 .697-.039.89a2 2 0 0 1-1.57 1.572C6.196 19 5.964 19 5.5 19s-.697 0-.89-.038a2 2 0 0 1-1.572-1.572C3 17.197 3 16.965 3 16.5v-2c0-.465 0-.697.038-.89a2 2 0 0 1 1.572-1.572C4.803 12 5.035 12 5.5 12s.697 0 .89.038a2 2 0 0 1 1.571 1.571c.039.194.039.426.039.891Z"/>`),
		g.Group(children),
	)
}