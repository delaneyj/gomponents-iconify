package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiskOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M12.778 17.012c0-.559.453-1.012 1.012-1.012h21.976c.559 0 1.012.453 1.012 1.012V32c0 6.627-5.373 12-12 12v0c-6.628 0-12-5.373-12-12V17.012Z"/><path d="M15.778 4h18v12h-18z"/><path stroke-linecap="round" d="M21.778 9v2m6-2v2m-15 21h24"/></g>`),
		g.Group(children),
	)
}