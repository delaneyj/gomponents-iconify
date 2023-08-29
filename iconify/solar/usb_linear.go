package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="20" r="2" stroke="currentColor" stroke-width="1.5"/><circle cx="6" cy="6" r="1" stroke="currentColor" stroke-width="1.5"/><path fill="currentColor" d="M11.25 18a.75.75 0 0 0 1.5 0h-1.5ZM12 2l.624-.416a.75.75 0 0 0-1.248 0L12 2Zm1.376 3.416a.75.75 0 1 0 1.248-.832l-1.248.832Zm-4-.832a.75.75 0 0 0 1.248.832l-1.248-.832Zm2-2.168l2 3l1.248-.832l-2-3l-1.248.832Zm0-.832l-2 3l1.248.832l2-3l-1.248-.832ZM12.75 18v-3h-1.5v3h1.5Zm0-3V2h-1.5v13h1.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M6 7v2.558a2 2 0 0 0 1.368 1.898l3.264 1.088A2 2 0 0 1 12 14.441V15m6-5v2.559a2 2 0 0 1-1.367 1.897l-3.265 1.088A2 2 0 0 0 12 17.442V18"/><path stroke="currentColor" stroke-width="1.5" d="M17 9c0-.471 0-.707.146-.854C17.293 8 17.53 8 18 8s.707 0 .854.146C19 8.293 19 8.53 19 9s0 .707-.146.854C18.707 10 18.47 10 18 10s-.707 0-.854-.146C17 9.707 17 9.47 17 9Z"/></g>`),
		g.Group(children),
	)
}