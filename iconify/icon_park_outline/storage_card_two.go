package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageCardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="32" height="40" x="8" y="4" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 4v10h20V4M14 24h20v20H14zm0 8h20m0-3v6m-20-6v6"/></g>`),
		g.Group(children),
	)
}