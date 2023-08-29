package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PercentRotateOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.168 8A10.002 10.002 0 0 0 12 2c-5.185 0-9.449 3.947-9.95 9"/><path d="M18 8h3.4a.6.6 0 0 0 .6-.6V4M2.881 16c1.544 3.532 5.068 6 9.168 6c5.186 0 9.45-3.947 9.951-9"/><path d="M6.05 16h-3.4a.6.6 0 0 0-.6.6V20"/><path fill="currentColor" d="M14.5 15a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-5-5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/><path d="m15 9l-6 6"/></g>`),
		g.Group(children),
	)
}