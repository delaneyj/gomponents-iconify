package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m12 24l-8-4m4 2v-4m32 4v-4"/><circle cx="24" cy="10" r="6" stroke="currentColor" stroke-width="4"/><ellipse cx="24" cy="30" stroke="currentColor" stroke-width="4" rx="13" ry="14"/><circle cx="24" cy="26" r="2" fill="currentColor"/><circle cx="24" cy="31" r="2" fill="currentColor"/><circle cx="24" cy="36" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m44 20l-8 4"/></g>`),
		g.Group(children),
	)
}