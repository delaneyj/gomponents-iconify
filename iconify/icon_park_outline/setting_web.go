package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingWeb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 40H7a3 3 0 0 1-3-3V11a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v12.059"/><path stroke="currentColor" stroke-width="4" d="M4 11a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v9H4v-9Z"/><circle r="2" fill="currentColor" transform="matrix(0 -1 -1 0 10 14)"/><circle r="2" fill="currentColor" transform="matrix(0 -1 -1 0 16 14)"/><circle cx="37" cy="34" r="3" stroke="currentColor" stroke-width="4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 41v-4m0-6v-4m-6.062 10.5l3.464-2m5.196-3l3.464-2m-12.124 0l3.464 2m5.196 3l3.464 2"/></g>`),
		g.Group(children),
	)
}