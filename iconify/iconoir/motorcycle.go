package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Motorcycle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm14-4l-3-9l1-1"/><path d="M16 8.5h-4.5l-4.5 3m-1.5 4H12l1-2.5l3.5-4.5m-8 1.5c-2-1.5-5-1.5-7 0M19 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g>`),
		g.Group(children),
	)
}