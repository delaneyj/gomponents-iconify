package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="30" height="24" x="9" y="18" stroke="currentColor" stroke-width="4" rx="2"/><circle cx="17" cy="26" r="2" fill="currentColor"/><circle cx="31" cy="26" r="2" fill="currentColor"/><path fill="currentColor" d="M20 32a2 2 0 1 0 0 4v-4Zm8 4a2 2 0 1 0 0-4v4Zm-8 0h8v-4h-8v4Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 10v8M4 26v8m40-8v8"/><circle cx="24" cy="8" r="2" stroke="currentColor" stroke-width="4"/></g>`),
		g.Group(children),
	)
}