package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FECB00" d="M.5.5h300v100H.5z"/><path fill="#EA2839" d="M.5 100.5h300v100H.5z"/><path fill="#34B233" d="M.5 67.166h300v66.667H.5z"/><path fill="#FFF" d="m80.407 84.759l43.32 31.473l-16.547 50.926l43.32-31.474l43.32 31.474l-16.547-50.926l43.32-31.473h-53.546L150.5 33.833l-16.546 50.926z"/></g>`),
		g.Group(children),
	)
}