package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 24h34M24 7v34"/><path fill="currentColor" d="M5 5h4v4H5zm9 0h4v4h-4zm16 0h4v4h-4zm9 0h4v4h-4zm0 9h4v4h-4zM5 14h4v4H5zm0 25h4v4H5zm9 0h4v4h-4zm16 0h4v4h-4zm9 0h4v4h-4zm0-9h4v4h-4zM5 30h4v4H5z"/></g>`),
		g.Group(children),
	)
}