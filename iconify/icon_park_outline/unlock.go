package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="34" height="22" x="7" y="22.048" rx="2"/><path stroke-linecap="round" d="M14 22v-7.995c-.005-5.135 3.923-9.438 9.086-9.954S32.967 6.974 34 12.006M24 30v6"/></g>`),
		g.Group(children),
	)
}