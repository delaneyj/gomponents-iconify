package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#00a04d" d="M0 360h640v120H0z"/><path fill="#151f6d" d="M0 120h640v120H0z"/><path fill="#ee2737" d="M0 0h640v120H0z"/><path fill="#ffcd00" d="M0 240h640v120H0z"/></g>`),
		g.Group(children),
	)
}