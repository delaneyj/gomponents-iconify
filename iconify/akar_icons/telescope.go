package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telescope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m7 21l6-8l8 8M5.228 7.303l13.532-5.06a1 1 0 0 1 1.108.285l.19.22A8 8 0 0 1 22 7.973v.616a1 1 0 0 1-.65.937l-13.536 5.06l-2.586-7.283Z"/><path d="M2.66 11.371a2 2 0 0 1 1.193-2.545l1.694-.624l1.944 5.473l-1.64.612a2 2 0 0 1-2.585-1.205l-.606-1.711ZM13 13v9"/></g>`),
		g.Group(children),
	)
}