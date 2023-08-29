package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.21 10.211c-3.405 3.852-6.975 5.279-10.71 4.281c3.711-.99 6.282-3.418 7.711-7.28L8.5 4.5h8v8z"/>`),
		g.Group(children),
	)
}