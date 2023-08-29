package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#009f4d" d="M0 384h512v128H0z"/><path fill="#151f6d" d="M0 128h512v128H0z"/><path fill="#ee2737" d="M0 0h512v128H0z"/><path fill="#ffcd00" d="M0 256h512v128H0z"/></g>`),
		g.Group(children),
	)
}