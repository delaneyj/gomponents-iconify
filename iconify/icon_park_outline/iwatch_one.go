package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IwatchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="22" height="24" x="13" y="12" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 12v6m-3-6h6m-3 18v6m-3 0h6m8-12.066L29 24m6-3v6m-16-3l-6 .066M13 21v6m18-15V4H17v8m14 24v7a1 1 0 0 1-1 1H18a1 1 0 0 1-1-1v-7M7 20v8m34-8v8"/></g>`),
		g.Group(children),
	)
}