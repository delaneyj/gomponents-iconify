package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Funds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><ellipse cx="14" cy="10" rx="10" ry="5"/><path d="M4 10v7c0 2.761 4.477 5 10 5s10-2.239 10-5v-7"/><path d="M4 17v7c0 2.761 4.477 5 10 5s10-2.239 10-5v-7"/><path d="M4 24v7c0 2.761 4.477 5 10 5s10-2.239 10-5v-7"/><path d="M4 31v7c0 2.761 4.477 5 10 5s10-2.239 10-5v-7"/><ellipse cx="34" cy="24" rx="10" ry="5"/><path d="M24 24v7c0 2.761 4.477 5 10 5s10-2.239 10-5v-7"/><path d="M24 31v7c0 2.761 4.477 5 10 5s10-2.239 10-5v-7"/></g>`),
		g.Group(children),
	)
}