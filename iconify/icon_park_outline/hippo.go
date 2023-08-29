package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hippo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="34" height="18" x="7" y="25" stroke="currentColor" stroke-linejoin="round" stroke-width="4" rx="6"/><circle cx="17" cy="34" r="3" stroke="currentColor" stroke-width="4"/><circle cx="31" cy="34" r="3" stroke="currentColor" stroke-width="4"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M11 19a6 6 0 0 1 6-6h14a6 6 0 0 1 6 6v6H11v-6Z"/><circle cx="20" cy="19" r="2" fill="currentColor"/><circle cx="28" cy="19" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M18 5a3 3 0 0 1 3 3v5h-6V8a3 3 0 0 1 3-3Zm12 0a3 3 0 0 1 3 3v5h-6V8a3 3 0 0 1 3-3Z"/></g>`),
		g.Group(children),
	)
}