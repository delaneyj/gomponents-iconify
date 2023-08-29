package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autoplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.194.665a8 8 0 1 0 4.194 10.404a.75.75 0 0 0-1.385-.575a6.5 6.5 0 1 1-.526-5.994H11.75a.75.75 0 0 0 0 1.5H16V1.75a.75.75 0 0 0-1.5 0v1.586a8 8 0 0 0-3.306-2.67Zm-.423 6.913a.5.5 0 0 1 0 .844l-4.003 2.54A.5.5 0 0 1 6 10.538V5.461a.5.5 0 0 1 .768-.422l4.003 2.539Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}